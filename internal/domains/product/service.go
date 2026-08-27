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
	sku, name, description, err := validateIdentityRequest(req.SKU, req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil, ErrInvalidCategoryID
	}

	brandID, err := uuid.Parse(req.BrandID)
	if err != nil {
		return nil, ErrInvalidBrandID
	}

	unitID, err := uuid.Parse(req.UnitID)
	if err != nil {
		return nil, ErrInvalidUnitID
	}

	if err := validateMeasurementRequest(
		req.MinStock,
		req.Weight,
		req.Length,
		req.Width,
	); err != nil {
		return nil, err
	}

	newProduct := &Product{
		SKU:         sku,
		Name:        name,
		CategoryID:  categoryID,
		BrandID:     brandID,
		UnitID:      unitID,
		MinStock:    req.MinStock,
		Weight:      req.Weight,
		Length:      req.Length,
		Width:       req.Width,
		Description: description,
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

	sku, name, description, err := validateIdentityRequest(req.SKU, req.Name, req.Description)
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

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil, ErrInvalidCategoryID
	}

	brandID, err := uuid.Parse(req.BrandID)
	if err != nil {
		return nil, ErrInvalidBrandID
	}

	unitID, err := uuid.Parse(req.UnitID)
	if err != nil {
		return nil, ErrInvalidUnitID
	}

	if err := validateMeasurementRequest(
		req.MinStock,
		req.Weight,
		req.Length,
		req.Width,
	); err != nil {
		return nil, err
	}

	newProduct := &Product{
		ID:          id,
		SKU:         sku,
		Name:        name,
		CategoryID:  categoryID,
		BrandID:     brandID,
		UnitID:      unitID,
		MinStock:    req.MinStock,
		Weight:      req.Weight,
		Length:      req.Length,
		Width:       req.Width,
		Description: description,
	}

	updatedProduct, err := s.repository.Update(ctx, newProduct)
	if err != nil {
		return nil, fmt.Errorf("update product: %w", err)
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

func validateIdentityRequest(rawSKU, rawName string, rawDescription *string) (string, string, *string, error) {
	sku := strings.TrimSpace(rawSKU)
	name := strings.TrimSpace(rawName)

	if sku == "" {
		return "", "", nil, ErrSKURequired
	}

	if name == "" {
		return "", "", nil, ErrNameRequired
	}

	if rawDescription == nil {
		return sku, name, nil, nil
	}

	description := strings.TrimSpace(*rawDescription)

	if description == "" {
		return sku, name, nil, nil
	}

	return sku, name, &description, nil
}

func validateMeasurementRequest(minStock int32, weight float32, length float32, width float32) error {
	if minStock < 0 {
		return ErrInvalidMinStock
	}

	if weight < 0 {
		return ErrInvalidWeight
	}

	if length < 0 {
		return ErrInvalidLength
	}

	if width < 0 {
		return ErrInvalidWidth
	}

	return nil
}
