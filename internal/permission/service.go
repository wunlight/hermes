package permission

import (
	"context"
	"fmt"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Permission, error) {
	existing, err := s.repository.GetByCode(ctx, req.Code)
	if err != nil {
		return nil, fmt.Errorf("get permission by code: %w", err)
	}

	if existing != nil {
		return nil, ErrPermissionAlreadyExists
	}

	permission, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create permission: %w", err)
	}

	return permission, nil
}
