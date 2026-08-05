package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/adapters/sqlc"
	"github.com/wunlight/hermes/internal/permission"
)

func Permission(router chi.Router, db sqlc.DBTX) {
	repo := permission.NewRepository(db)
	service := permission.NewService(repo)
	handler := permission.NewHandler(service)
	permission.Mount(router, handler)
}
