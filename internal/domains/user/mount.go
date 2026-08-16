package user

import (
	"github.com/go-chi/chi/v5"

	"github.com/wunlight/hermes/internal/middleware"
)

func Mount(r chi.Router, handler *Handler, auth *middleware.AuthMiddleware) {
	r.Route("/users", func(r chi.Router) {
		r.Use(auth.RequireAuth)

		r.Get("/{id}", handler.GetByID)
		r.Patch("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
		r.Get("/", handler.List)
	})
}
