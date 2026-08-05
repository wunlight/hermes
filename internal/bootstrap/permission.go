package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wunlight/hermes/internal/permission"
)

func Permission(router chi.Router, db *pgxpool.Pool) {
	repo := permission.NewRepository(db)
	service := permission.NewService(repo)
	handler := permission.NewHandler(service)
	permission.Mount(router, handler)
}
