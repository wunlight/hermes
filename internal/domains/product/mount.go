package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/middleware"
)

func Mount(r chi.Router, handler *Handler, auth *middleware.AuthMiddleware) {
	r.Route("/product", func(r chi.Router) {
		r.Use(auth.RequireAuth)

		r.Get("/", handler.List)
		r.Get("/{id}", handler.GetByID)
		r.Post("/", handler.Create)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}
