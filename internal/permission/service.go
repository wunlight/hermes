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
	permission, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create permission: %w", err)
	}

	return permission, nil
}
