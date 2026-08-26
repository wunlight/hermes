package category

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Test_List(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	categories, err := repo.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, category := range categories {
		t.Logf(
			"ID=%s ParentID=%s Code=%s Name=%s",
			category.ID,
			category.ParentID,
			category.Code,
			category.Name,
		)
	}
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
