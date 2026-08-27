package category

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

func (s *Service) List(ctx context.Context) ([]*Category, error) {
	result, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list category: %w", err)
	}

	return result, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Category, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	result, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get category by id: %w", err)
	}

	return result, nil
}

func (s *Service) Create(ctx context.Context, req CreateReq) (*Category, error) {
	var parentID *uuid.UUID

	if req.ParentID != nil {
		parsedParentID, err := s.validateParent(ctx, req.ParentID)
		if err != nil {
			return nil, err
		}

		parentID = parsedParentID
	}

	code, name, err := validateIdentityRequest(req.Code, req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.validateCodeUnique(ctx, code, nil); err != nil {
		return nil, err
	}

	newCategory := &Category{
		ParentID: parentID,
		Code:     code,
		Name:     name,
	}

	createdCategory, err := s.repository.Create(ctx, newCategory)
	if err != nil {
		return nil, fmt.Errorf("create category: %w", err)
	}

	return createdCategory, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateReq) (*Category, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	var parentID *uuid.UUID

	if req.ParentID != nil {
		parsedParentID, err := s.validateParent(ctx, req.ParentID)
		if err != nil {
			return nil, err
		}

		parentID = parsedParentID
	}

	code, name, err := validateIdentityRequest(req.Code, req.Name)
	if err != nil {
		return nil, err
	}

	_, err = s.repository.GetByID(ctx, id)
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("check existing category: %w", err)
		}

		return nil, ErrNotFound
	}

	if err := s.validateCodeUnique(ctx, code, &id); err != nil {
		return nil, err
	}

	newCategory := &Category{
		ID:       id,
		ParentID: parentID,
		Code:     code,
		Name:     name,
	}

	updatedCategory, err := s.repository.Update(ctx, newCategory)
	if err != nil {
		return nil, fmt.Errorf("update category: %w", err)
	}

	return updatedCategory, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	if _, err := s.repository.GetByID(ctx, id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrNotFound
		}

		return fmt.Errorf("get category: %w", err)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	return nil
}

func (s *Service) validateParent(ctx context.Context, id *string) (*uuid.UUID, error) {
	parsedParentID, err := uuid.Parse(*id)
	if err != nil {
		return nil, ErrInvalidParentID
	}

	_, err = s.repository.GetByID(ctx, parsedParentID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrInvalidParentID
		}

		return nil, fmt.Errorf("check parent category: %w", err)
	}

	return &parsedParentID, nil
}

func (s *Service) validateCodeUnique(ctx context.Context, code string, excludeID *uuid.UUID) error {
	existingCategory, err := s.repository.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}

		return fmt.Errorf("check existing category: %w", err)
	}

	if excludeID != nil && existingCategory.ID == *excludeID {
		return nil
	}

	return ErrCategoryAlreadyExists
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
