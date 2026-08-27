package unit

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	units, err := h.service.List(r.Context())
	if err != nil {
		h.handleError(w, err)
		return
	}

	res := make([]*UnitResponse, 0, len(units))
	for _, unit := range units {
		res = append(res, toResponse(unit))
	}

	writeJSON(w, http.StatusOK, res)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, ErrInvalidID.Error())
		return
	}

	unit, err := h.service.GetByID(r.Context(), id)

	writeJSON(w, http.StatusOK, toResponse(unit))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, ErrInvalidReq.Error())
		return
	}

	unit, err := h.service.Create(r.Context(), req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toResponse(unit))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, ErrInvalidID.Error())
		return
	}

	var req UpdateReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, ErrInvalidReq.Error())
		return
	}

	unit, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toResponse(unit))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, ErrInvalidID.Error())
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

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
	case errors.Is(err, ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())

	case errors.Is(err, ErrInvalidID):
		writeError(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, ErrNameRequired),
		errors.Is(err, ErrCodeRequired):
		writeError(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, ErrCodeAlreadyExists):
		writeError(w, http.StatusConflict, err.Error())

	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func toResponse(category *Unit) *UnitResponse {
	if category == nil {
		return nil
	}

	return &UnitResponse{
		ID:   category.ID,
		Code: category.Code,
		Name: category.Name,
	}
}
