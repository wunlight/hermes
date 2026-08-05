package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wunlight/hermes/internal/role"
)

func Role(router chi.Router, db *pgxpool.Pool) {
	repo := role.NewRepository(db)
	service := role.NewService(repo)
	handler := role.NewHandler(service)
	role.Mount(router, handler)
}
