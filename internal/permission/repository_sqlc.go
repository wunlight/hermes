package permission

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
)

type sqlcRepository struct {
	queries *sqlc.Queries
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &sqlcRepository{
		queries: sqlc.New(pool),
	}
}

func (r *sqlcRepository) Create(ctx context.Context, req CreateRequest) (*Permission, error) {
	row, err := r.queries.CreatePermission(
		ctx,
		sqlc.CreatePermissionParams{
			Code: req.Code,
			Description: pgtype.Text{
				String: req.Description,
				Valid:  req.Description != "",
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create permission: %w", err)
	}

	return toDomain(row), nil
}

func toDomain(p sqlc.Permission) *Permission {
	return &Permission{
		ID:          p.ID,
		Code:        p.Code,
		Description: p.Description.String,
		CreatedAt:   p.CreatedAt.Time,
		UpdatedAt:   p.UpdatedAt.Time,
	}
}
