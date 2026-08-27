package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
	"github.com/wunlight/hermes/internal/infrastructure/adapters/pg_util"
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

	weight, length, width, err := measurementsToNumeric(product.Weight, product.Length, product.Width)
	if err != nil {
		return nil, fmt.Errorf("convert measurements: %w", err)
	}

	row, err := queries.CreateProduct(
		ctx,
		sqlc.CreateProductParams{
			Sku:         product.SKU,
			Name:        product.Name,
			CategoryID:  product.CategoryID,
			BrandID:     &product.BrandID,
			UnitID:      product.UnitID,
			MinStock:    product.MinStock,
			Weight:      weight,
			Length:      length,
			Width:       width,
			Description: pg_util.NullableStringToText(product.Description),
		},
	)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, ErrSKUAlreadyExists
		}

		return nil, fmt.Errorf("create product: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, product *Product) (*Product, error) {
	queries := sqlc.New(r.db)

	weight, length, width, err := measurementsToNumeric(product.Weight, product.Length, product.Width)
	if err != nil {
		return nil, fmt.Errorf("convert measurements: %w", err)
	}

	row, err := queries.UpdateProduct(
		ctx,
		sqlc.UpdateProductParams{
			ID:          product.ID,
			Sku:         product.SKU,
			Name:        product.Name,
			CategoryID:  product.CategoryID,
			BrandID:     &product.BrandID,
			UnitID:      product.UnitID,
			MinStock:    product.MinStock,
			Weight:      weight,
			Length:      length,
			Width:       width,
			Description: pg_util.NullableStringToText(product.Description),
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

func measurementsToNumeric(weight float32, length float32, width float32) (pgtype.Numeric, pgtype.Numeric, pgtype.Numeric, error) {
	w, err := pg_util.MeasurementToNumeric(weight)
	if err != nil {
		return pgtype.Numeric{}, pgtype.Numeric{}, pgtype.Numeric{}, err
	}

	l, err := pg_util.MeasurementToNumeric(length)
	if err != nil {
		return pgtype.Numeric{}, pgtype.Numeric{}, pgtype.Numeric{}, err
	}

	wd, err := pg_util.MeasurementToNumeric(width)
	if err != nil {
		return pgtype.Numeric{}, pgtype.Numeric{}, pgtype.Numeric{}, err
	}

	return w, l, wd, nil
}
