package brand

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) List(ctx context.Context) ([]*Brand, error) {
	res, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list brand: %w", err)
	}

	return res, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Brand, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	res, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get brand by id: %w", err)
	}

	return res, nil
}

func (s *Service) Create(ctx context.Context, req CreateReq) (*Brand, error) {
	code, name, err := validateIdentityRequest(req.Code, req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.validateCodeUnique(ctx, code, nil); err != nil {
		return nil, err
	}

	newBrand := &Brand{
		Code: code,
		Name: name,
	}

	createdBrand, err := s.repository.Create(ctx, newBrand)
	if err != nil {
		return nil, fmt.Errorf("create brand: %w", err)
	}

	return createdBrand, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateReq) (*Brand, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	code, name, err := validateIdentityRequest(req.Code, req.Name)
	if err != nil {
		return nil, err
	}

	_, err = s.GetByID(ctx, id)
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("check existing brand: %w", err)
		}

		return nil, ErrNotFound
	}

	if err := s.validateCodeUnique(ctx, code, &id); err != nil {
		return nil, err
	}

	newBrand := &Brand{
		ID:   id,
		Code: code,
		Name: name,
	}

	updatedBrand, err := s.repository.Update(ctx, newBrand)
	if err != nil {
		return nil, fmt.Errorf("create brand: %w", err)
	}

	return updatedBrand, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return fmt.Errorf("check existing brand: %w", err)
		}

		return ErrNotFound
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete brand: %w", err)
	}

	return nil
}

func (s *Service) validateCodeUnique(ctx context.Context, code string, excludeID *uuid.UUID) error {
	existingBrand, err := s.repository.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}

		return fmt.Errorf("check existing brand: %w", err)
	}

	if excludeID != nil && existingBrand.ID == *excludeID {
		return nil
	}

	return ErrCodeAlreadyExists
}

func validateIdentityRequest(rawCode, rawName string) (string, string, error) {
	code := strings.TrimSpace(rawCode)
	name := strings.TrimSpace(rawName)

	if code == "" {
		return "", "", ErrCodeRequired
	}

	if name == "" {
		return "", "", ErrNameRequired
	}

	return code, name, nil
}
