package role

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
	return &sqlcRepository{queries: sqlc.New(pool)}
}

func (r *sqlcRepository) Create(ctx context.Context, req CreateRequest) (*Role, error) {
	row, err := r.queries.CreateRole(
		ctx,
		sqlc.CreateRoleParams{
			Code: req.Code,
			Name: req.Name,
			Description: pgtype.Text{
				String: req.Description,
				Valid:  req.Description != "",
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	return toDomain(row), nil
}

func toDomain(p sqlc.Role) *Role {
	return &Role{
		ID:          p.ID,
		Code:        p.Code,
		Name:        p.Name,
		Description: p.Description.String,
		CreatedAt:   p.CreatedAt.Time,
		UpdatedAt:   p.UpdatedAt.Time,
	}
}
