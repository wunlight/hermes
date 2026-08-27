package brand

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	List(ctx context.Context) ([]*Brand, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Brand, error)
	GetByCode(ctx context.Context, code string) (*Brand, error)
	Create(ctx context.Context, brand *Brand) (*Brand, error)
	Update(ctx context.Context, brand *Brand) (*Brand, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
