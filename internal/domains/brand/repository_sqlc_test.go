package brand_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wunlight/hermes/internal/domains/brand"
)

func Test_List(t *testing.T) {
	db := setupTestDB(t)
	repo := brand.NewRepository(db)

	brands, err := repo.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("brands: %d", len(brands))
	t.Logf("err: %s", err)

	for _, brand := range brands {
		t.Logf(
			"ID=%s Code=%s Name=%s",
			brand.ID,
			brand.Code,
			brand.Name,
		)
	}
}

func Test_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := brand.NewRepository(db)

	brand, err := repo.GetByID(context.Background(), uuid.Nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("brand: %s, %s, %s", brand.ID, brand.Code, brand.Name)
	t.Logf("err: %s", err)
}

func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "host=localhost port=5432 user=argus password=4rGu5! dbname=atlas sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}
