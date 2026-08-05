package role

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

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Role, error) {
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return nil, ErrRoleCodeRequired
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, ErrRoleNameRequired
	}

	role, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	return role, nil
}
