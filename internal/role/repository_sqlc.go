package role

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
)

type sqlcRepository struct {
	queries *sqlc.Queries
}

func NewRepository(db sqlc.DBTX) Repository {
	return &sqlcRepository{queries: sqlc.New(db)}
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

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Role, error) {
	row, err := r.queries.GetRoleByCode(
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

func (r *sqlcRepository) CreateRolePermission(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	return r.queries.CreateRolePermission(
		ctx,
		sqlc.CreateRolePermissionParams{
			RoleID:       roleID,
			PermissionID: permissionID,
		},
	)
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
