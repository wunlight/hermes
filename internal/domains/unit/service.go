package unit

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

func (s *Service) List(ctx context.Context) ([]*Unit, error) {
	res, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list unit: %w", err)
	}

	return res, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Unit, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	res, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get unit by id: %w", err)
	}

	return res, nil
}

func (s *Service) Create(ctx context.Context, req CreateReq) (*Unit, error) {
	code, name, err := validateIdentityRequest(req.Code, req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.validateCodeUnique(ctx, code, nil); err != nil {
		return nil, err
	}

	newUnit := &Unit{
		Code: code,
		Name: name,
	}

	createdUnit, err := s.repository.Create(ctx, newUnit)
	if err != nil {
		return nil, fmt.Errorf("create unit: %w", err)
	}

	return createdUnit, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateReq) (*Unit, error) {
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
			return nil, fmt.Errorf("check existing unit: %w", err)
		}

		return nil, ErrNotFound
	}

	if err := s.validateCodeUnique(ctx, code, &id); err != nil {
		return nil, err
	}

	newUnit := &Unit{
		ID:   id,
		Code: code,
		Name: name,
	}

	updatedUnit, err := s.repository.Update(ctx, newUnit)
	if err != nil {
		return nil, fmt.Errorf("create unit: %w", err)
	}

	return updatedUnit, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return fmt.Errorf("check existing unit: %w", err)
		}

		return ErrNotFound
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete unit: %w", err)
	}

	return nil
}

func (s *Service) validateCodeUnique(ctx context.Context, code string, excludeID *uuid.UUID) error {
	existingUnit, err := s.repository.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}

		return fmt.Errorf("check existing unit: %w", err)
	}

	if excludeID != nil && existingUnit.ID == *excludeID {
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
