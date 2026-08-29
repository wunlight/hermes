package category

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

func (r *sqlcRepository) List(ctx context.Context) ([]*Category, error) {
	queries := sqlc.New(r.db)

	rows, err := queries.ListCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("list categories: %w", err)
	}

	categories := make([]*Category, 0, len(rows))

	for _, row := range rows {
		categories = append(categories, toDomain(row))
	}

	return categories, nil
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*Category, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetCategoryByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get category by id: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByCode(ctx context.Context, code string) (*Category, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetCategoryByCode(ctx, code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get category by code: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Create(ctx context.Context, category *Category) (*Category, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateCategory(
		ctx,
		sqlc.CreateCategoryParams{
			ParentID: category.ParentID,
			Code:     category.Code,
			Name:     category.Name,
		},
	)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, ErrCodeAlreadyExists
		}

		return nil, fmt.Errorf("create category: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, category *Category) (*Category, error) {
	queries := sqlc.New(r.db)

	row, err := queries.UpdateCategory(
		ctx,
		sqlc.UpdateCategoryParams{
			ID:       category.ID,
			ParentID: category.ParentID,
			Code:     category.Code,
			Name:     category.Name,
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("update category: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	queries := sqlc.New(r.db)

	result, err := queries.DeleteCategory(ctx, id)
	if err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func toDomain(row sqlc.Category) *Category {
	return &Category{
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
