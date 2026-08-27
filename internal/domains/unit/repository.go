package unit

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	List(ctx context.Context) ([]*Unit, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Unit, error)
	GetByCode(ctx context.Context, code string) (*Unit, error)
	Create(ctx context.Context, unit *Unit) (*Unit, error)
	Update(ctx context.Context, unit *Unit) (*Unit, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
