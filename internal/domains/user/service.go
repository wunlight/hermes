package user

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	result, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return result, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (*User, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	name := strings.TrimSpace(req.Name)

	if err := validateUpdateRequest(email, name); err != nil {
		return nil, err
	}

	currentUser, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get user: %w", err)
	}

	if email != currentUser.Email {
		existingUser, err := s.repository.GetByEmail(ctx, email)
		if err == nil && existingUser.ID != id {
			return nil, ErrEmailAlreadyExists
		}

		if err != nil && !errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("check email: %w", err)
		}
	}

	currentUser.Email = email
	currentUser.Name = name

	updatedUser, err := s.repository.Update(ctx, currentUser)
	if err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	return updatedUser, nil
}

func validateUpdateRequest(email string, name string) error {
	if email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	if name == "" {
		return ErrNameRequired
	}

	if len(name) > 255 {
		return ErrNameTooLong
	}

	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	if _, err := s.repository.GetByID(ctx, id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrNotFound
		}

		return fmt.Errorf("get user: %w", err)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (s *Service) List(ctx context.Context) ([]*User, error) {
	users, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	return users, nil
}
