package brand

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// func (h *Handler) List(w http.ResponseWriter, r *http.Request)

// func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request)

// func (h *Handler) Create(w http.ResponseWriter, r *http.Request)

// func (h *Handler) Update(w http.ResponseWriter, r *http.Request)

// func (h *Handler) Delete(w http.ResponseWriter, r *http.Request)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	switch {
	// case errors.Is(err, ErrNotFound):
	// 	writeError(w, http.StatusNotFound, "category not found")

	// case errors.Is(err, ErrInvalidID):
	// 	writeError(w, http.StatusBadRequest, "invalid category id")

	// case errors.Is(err, ErrNameRequired),
	// 	errors.Is(err, ErrCodeRequired):
	// 	writeError(w, http.StatusBadRequest, err.Error())

	// case errors.Is(err, ErrCategoryAlreadyExists):
	// 	writeError(w, http.StatusConflict, err.Error())

	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func toResponse(category *Brand) *BrandResponse {
	if category == nil {
		return nil
	}

	return &BrandResponse{
		ID:        category.ID,
		Code:      category.Code,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
