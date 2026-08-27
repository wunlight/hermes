package unit

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
)

type sqlcRepository struct {
	db sqlc.DBTX
}

func NewRepository(db sqlc.DBTX) Repository {
	return &sqlcRepository{db: db}
}

func (r *sqlcRepository) List(ctx context.Context) ([]*Unit, error) {
	queries := sqlc.New(r.db)

	rows, err := queries.ListUnits(ctx)
	if err != nil {
		return nil, fmt.Errorf("list units: %w", err)
	}

	units := make([]*Unit, 0, len(rows))

	for _, row := range rows {
		units = append(units, toDomain(row))
	}

	return units, nil
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*Unit, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetUnitByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get unit by id: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Unit, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetUnitByCode(ctx, code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get unit by code: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Create(ctx context.Context, unit *Unit) (*Unit, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateUnit(
		ctx,
		sqlc.CreateUnitParams{
			Code: unit.Code,
			Name: unit.Name,
		},
	)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, ErrCodeAlreadyExists
		}

		return nil, fmt.Errorf("create unit: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, unit *Unit) (*Unit, error) {
	queries := sqlc.New(r.db)

	row, err := queries.UpdateUnit(
		ctx,
		sqlc.UpdateUnitParams{
			ID:   unit.ID,
			Code: unit.Code,
			Name: unit.Name,
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("update unit: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	queries := sqlc.New(r.db)

	if err := queries.DeleteUnit(ctx, id); err != nil {
		return fmt.Errorf("delete unit: %w", err)
	}

	return nil
}

func toDomain(row sqlc.Unit) *Unit {
	return &Unit{
		ID:   row.ID,
		Code: row.Code,
		Name: row.Name,
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) &&
		pgErr.Code == "23505"
}
