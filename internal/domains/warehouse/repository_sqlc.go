package warehouse

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
	"github.com/wunlight/hermes/internal/infrastructure/adapters/pg_util"
)

type sqlcRepository struct {
	db sqlc.DBTX
}

func NewRepository(db sqlc.DBTX) Repository {
	return &sqlcRepository{db: db}
}

func (r *sqlcRepository) List(ctx context.Context) ([]*Warehouse, error) {
	queries := sqlc.New(r.db)

	rows, err := queries.ListWarehouses(ctx)
	if err != nil {
		return nil, fmt.Errorf("list warehouses: %w", err)
	}

	warehouses := make([]*Warehouse, 0, len(rows))

	for _, row := range rows {
		warehouses = append(warehouses, toDomain(row))
	}

	return warehouses, nil
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*Warehouse, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetWarehouseByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get warehouse by id: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Warehouse, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetWarehouseByCode(ctx, code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get warehouse by code: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Create(ctx context.Context, warehouse *Warehouse) (*Warehouse, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateWarehouse(
		ctx,
		sqlc.CreateWarehouseParams{
			Code:        warehouse.Code,
			Name:        warehouse.Name,
			Description: pg_util.NullableStringToText(warehouse.Description),
		},
	)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, ErrCodeAlreadyExists
		}

		return nil, fmt.Errorf("create warehouse: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, warehouse *Warehouse) (*Warehouse, error) {
	queries := sqlc.New(r.db)

	row, err := queries.UpdateWarehouse(
		ctx,
		sqlc.UpdateWarehouseParams{
			ID:          warehouse.ID,
			Code:        warehouse.Code,
			Name:        warehouse.Name,
			Description: pg_util.NullableStringToText(warehouse.Description),
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("update warehouse: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	queries := sqlc.New(r.db)

	result, err := queries.DeleteWarehouse(ctx, id)
	if err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func toDomain(row sqlc.Warehouse) *Warehouse {
	return &Warehouse{
		ID:          row.ID,
		Code:        row.Code,
		Name:        row.Name,
		Description: pg_util.TextToNullableString(row.Description),
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) &&
		pgErr.Code == "23505"
}
