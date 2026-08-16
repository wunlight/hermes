package auth

import "github.com/go-chi/chi/v5"

func Mount(r chi.Router, handler *Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.Refresh)
		r.Post("/logout", handler.Logout)
	})
}
