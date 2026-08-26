package brand

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) List(ctx context.Context) ([]*Brand, error)

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Brand, error)

func (s *Service) Create(ctx context.Context, category *Brand) (*Brand, error)

func (s *Service) Update(ctx context.Context, category *Brand) (*Brand, error)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error
