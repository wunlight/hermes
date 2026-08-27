package product

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

func (s *Service) List(ctx context.Context) ([]*Product, error) {
	res, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list product: %w", err)
	}

	return res, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Product, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	res, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get product by id: %w", err)
	}

	return res, nil
}

func (s *Service) Create(ctx context.Context, req CreateReq) (*Product, error) {
	code, name, err := validateIdentityRequest(req.Code, req.Name)
	if err != nil {
		return nil, err
	}

	newProduct := &Product{
		Code: code,
		Name: name,
	}

	createdProduct, err := s.repository.Create(ctx, newProduct)
	if err != nil {
		return nil, fmt.Errorf("create product: %w", err)
	}

	return createdProduct, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateReq) (*Product, error) {
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
			return nil, fmt.Errorf("check existing product: %w", err)
		}

		return nil, ErrNotFound
	}

	newProduct := &Product{
		ID:   id,
		Code: code,
		Name: name,
	}

	updatedProduct, err := s.repository.Update(ctx, newProduct)
	if err != nil {
		return nil, fmt.Errorf("create product: %w", err)
	}

	return updatedProduct, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return fmt.Errorf("check existing product: %w", err)
		}

		return ErrNotFound
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete product: %w", err)
	}

	return nil
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
