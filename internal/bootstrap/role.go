package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wunlight/hermes/internal/permission"
	"github.com/wunlight/hermes/internal/role"
)

func Role(router chi.Router, db *pgxpool.Pool) {
	repo := role.NewRepository(db)
	permissionRepo := permission.NewRepository(db)
	service := role.NewService(repo, permissionRepo)
	handler := role.NewHandler(service)
	role.Mount(router, handler)
}
