package product

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

func (r *sqlcRepository) List(ctx context.Context) ([]*Product, error) {
	queries := sqlc.New(r.db)

	rows, err := queries.ListProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("list products: %w", err)
	}

	products := make([]*Product, 0, len(rows))

	for _, row := range rows {
		products = append(products, toDomain(row))
	}

	return products, nil
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*Product, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get product by id: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Create(ctx context.Context, product *Product) (*Product, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateProduct(
		ctx,
		sqlc.CreateProductParams{
			Name: product.Name,
		},
	)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, ErrCodeAlreadyExists
		}

		return nil, fmt.Errorf("create product: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, product *Product) (*Product, error) {
	queries := sqlc.New(r.db)

	row, err := queries.UpdateProduct(
		ctx,
		sqlc.UpdateProductParams{
			ID:   product.ID,
			Name: product.Name,
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("update product: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	queries := sqlc.New(r.db)

	if err := queries.DeleteProduct(ctx, id); err != nil {
		return fmt.Errorf("delete product: %w", err)
	}

	return nil
}

func toDomain(row sqlc.Product) *Product {
	return &Product{
		ID:   row.ID,
		Name: row.Name,
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) &&
		pgErr.Code == "23505"
}
