package role

import "github.com/go-chi/chi/v5"

func Mount(r chi.Router, handler *Handler) {
	r.Route("/roles", func(r chi.Router) {
		r.Post("/", handler.Create)
	})
}
