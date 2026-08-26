package brand

import (
	"context"

	"github.com/google/uuid"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
	"github.com/wunlight/hermes/internal/infrastructure/adapters/pg_util"
)

type sqlcRepository struct {
	db sqlc.DBTX
}

func NewRepository(db sqlc.DBTX) Repository {
	return &sqlcRepository{db: db}
}

func (r *sqlcRepository) List(ctx context.Context) ([]*Brand, error)

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*Brand, error)

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Brand, error)

func (r *sqlcRepository) Create(ctx context.Context, brand *Brand) (*Brand, error)

func (r *sqlcRepository) Update(ctx context.Context, brand *Brand) (*Brand, error)

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error

func toDomain(row sqlc.Brand) *Brand {
	return &Brand{
		ID:        row.ID,
		Code:      row.Code,
		Name:      row.Name,
		CreatedAt: *pg_util.TimestamptzToTime(row.CreatedAt),
		UpdatedAt: *pg_util.TimestamptzToTime(row.UpdatedAt),
		DeletedAt: pg_util.TimestamptzToTime(row.DeletedAt),
	}
}
