package warehouse

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	List(ctx context.Context) ([]*Warehouse, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Warehouse, error)
	GetByCode(ctx context.Context, code string) (*Warehouse, error)
	Create(ctx context.Context, warehouse *Warehouse) (*Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) (*Warehouse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
