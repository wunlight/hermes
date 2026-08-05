package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/wunlight/hermes/internal/permission"
)

type Service struct {
	repository           Repository
	permissionRepository permission.Repository
}

func NewService(repository Repository, permissionRepository permission.Repository) *Service {
	return &Service{
		repository:           repository,
		permissionRepository: permissionRepository,
	}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Role, error) {
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return nil, ErrRoleCodeRequired
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, ErrRoleNameRequired
	}

	normalized, err := validatePermissionCodes(req.PermissionCodes)
	if err != nil {
		return nil, err
	}
	req.PermissionCodes = normalized

	permissions, err := s.permissionRepository.GetByCodes(
		ctx,
		req.PermissionCodes,
	)
	if err != nil {
		return nil, fmt.Errorf("get permissions: %w", err)
	}
	if len(permissions) != len(req.PermissionCodes) {
		return nil, ErrPermissionNotFound
	}

	role, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	return role, nil
}
