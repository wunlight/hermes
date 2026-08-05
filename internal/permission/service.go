package permission

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Permission, error) {
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return nil, ErrPermissionCodeRequired
	}

	req.Code = strings.ToLower(req.Code)
	if !isValidPermissionCode(req.Code) {
		return nil, ErrInvalidPermissionCode
	}

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
