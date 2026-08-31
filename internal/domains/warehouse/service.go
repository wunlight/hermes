package warehouse

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

func (s *Service) List(ctx context.Context) ([]*Warehouse, error) {
	res, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list warehouse: %w", err)
	}

	return res, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Warehouse, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	res, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get warehouse by id: %w", err)
	}

	return res, nil
}

func (s *Service) Create(ctx context.Context, req CreateReq) (*Warehouse, error) {
	code, name, description, err := validateIdentityRequest(req.Code, req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	if err := s.validateCodeUnique(ctx, code, nil); err != nil {
		return nil, err
	}

	newWarehouse := &Warehouse{
		Code:        code,
		Name:        name,
		Description: description,
	}

	createdWarehouse, err := s.repository.Create(ctx, newWarehouse)
	if err != nil {
		return nil, fmt.Errorf("create warehouse: %w", err)
	}

	return createdWarehouse, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateReq) (*Warehouse, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	code, name, description, err := validateIdentityRequest(req.Code, req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	_, err = s.GetByID(ctx, id)
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("check existing warehouse: %w", err)
		}

		return nil, ErrNotFound
	}

	if err := s.validateCodeUnique(ctx, code, &id); err != nil {
		return nil, err
	}

	newWarehouse := &Warehouse{
		ID:          id,
		Code:        code,
		Name:        name,
		Description: description,
	}

	updatedWarehouse, err := s.repository.Update(ctx, newWarehouse)
	if err != nil {
		return nil, fmt.Errorf("create warehouse: %w", err)
	}

	return updatedWarehouse, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return fmt.Errorf("check existing warehouse: %w", err)
		}

		return ErrNotFound
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete warehouse: %w", err)
	}

	return nil
}

func (s *Service) validateCodeUnique(ctx context.Context, code string, excludeID *uuid.UUID) error {
	existingWarehouse, err := s.repository.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}

		return fmt.Errorf("check existing warehouse: %w", err)
	}

	if excludeID != nil && existingWarehouse.ID == *excludeID {
		return nil
	}

	return ErrCodeAlreadyExists
}

func validateIdentityRequest(rawCode, rawName string, rawDesc *string) (string, string, *string, error) {
	code := strings.TrimSpace(rawCode)
	name := strings.TrimSpace(rawName)

	if code == "" {
		return "", "", nil, ErrCodeRequired
	}

	if name == "" {
		return "", "", nil, ErrNameRequired
	}

	var desc *string

	if rawDesc != nil {
		value := strings.TrimSpace(*rawDesc)

		if value != "" {
			desc = &value
		}
	}

	return code, name, desc, nil
}
