package permission

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	permission, err := h.service.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, ErrPermissionCodeRequired) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors.Is(err, ErrInvalidPermissionCode) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors.Is(err, ErrPermissionAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(toResponse(permission)); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func toResponse(p *Permission) *Response {
	return &Response{
		ID:          p.ID,
		Code:        p.Code,
		Description: p.Description,
	}
}
