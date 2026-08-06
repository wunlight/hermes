package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	return r
}

// func Register(r chi.Router, db *database.DB) {
// 	r.Get("/health", health)

// 	bootstrap.Permission(r, db)
// 	bootstrap.Role(r, db)
// }

// func health(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	_, _ = w.Write([]byte(`{"status":"ok"}`))
// }
