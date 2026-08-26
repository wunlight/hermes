package brand

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

func (r *sqlcRepository) List(ctx context.Context) ([]*Brand, error) {
	queries := sqlc.New(r.db)

	rows, err := queries.ListBrands(ctx)
	if err != nil {
		return nil, fmt.Errorf("list brands: %w", err)
	}

	brands := make([]*Brand, 0, len(rows))

	for _, row := range rows {
		brands = append(brands, toDomain(row))
	}

	return brands, nil
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*Brand, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetBrandByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get brand by id: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Brand, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetBrandByCode(ctx, code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get brand by code: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Create(ctx context.Context, brand *Brand) (*Brand, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateBrand(
		ctx,
		sqlc.CreateBrandParams{
			Code: brand.Code,
			Name: brand.Name,
		},
	)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, ErrCodeAlreadyExists
		}

		return nil, fmt.Errorf("create brand: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, brand *Brand) (*Brand, error) {
	queries := sqlc.New(r.db)

	row, err := queries.UpdateBrand(
		ctx,
		sqlc.UpdateBrandParams{
			ID:   brand.ID,
			Code: brand.Code,
			Name: brand.Name,
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("update brand: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	queries := sqlc.New(r.db)

	result, err := queries.DeleteBrand(ctx, id)
	if err != nil {
		return fmt.Errorf("delete brand: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

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

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) &&
		pgErr.Code == "23505"
}
