package permission

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Permission, error) {
	row, err := r.queries.GetPermissionByCode(
		ctx,
		code,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get permission by code: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByCodes(ctx context.Context, codes []string) ([]Permission, error) {
	rows, err := r.queries.GetPermissionByCodes(
		ctx,
		codes,
	)
	if err != nil {
		return nil, fmt.Errorf("get permission by codes: %w", err)
	}

	return toDomains(rows), nil
}

func toDomains(rows []sqlc.Permission) []Permission {
	permissions := make([]Permission, 0, len(rows))

	for _, row := range rows {
		permissions = append(permissions, *toDomain(row))
	}

	return permissions
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
