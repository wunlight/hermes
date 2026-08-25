package category

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	GetByCode(ctx context.Context, code string) (*Category, error)
	Update(ctx context.Context, category *Category) (*Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*Category, error)
}
