package role

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	roleRepository       Repository
	permissionRepository PermissionLookup
}

func NewService(repository Repository, permissionRepository PermissionLookup) *Service {
	return &Service{
		roleRepository:       repository,
		permissionRepository: permissionRepository,
	}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Role, error) {
	req.Code = strings.ToLower(
		strings.TrimSpace(req.Code),
	)
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

	existing, err := s.roleRepository.GetByCode(ctx, req.Code)
	if err != nil {
		return nil, fmt.Errorf("get role by code: %w", err)
	}
	if existing != nil {
		return nil, ErrRoleAlreadyExists
	}

	newRole, err := s.roleRepository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	for _, permission := range permissions {
		if err := s.roleRepository.CreateRolePermission(
			ctx,
			newRole.ID,
			permission.ID,
		); err != nil {
			return nil, fmt.Errorf("assign permission: %w", err)
		}
	}

	return newRole, nil
}
