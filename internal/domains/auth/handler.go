package auth

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	service             *Service
	refreshCookieName   string
	refreshCookieMaxAge int
	secureCookie        bool
}

func NewHandler(
	service *Service,
	refreshCookieName string,
	refreshCookieMaxAge int,
	secureCookie bool,
) *Handler {
	return &Handler{
		service:             service,
		refreshCookieName:   refreshCookieName,
		refreshCookieMaxAge: refreshCookieMaxAge,
		secureCookie:        secureCookie,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.service.Register(r.Context(), req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	response := RegisterResponse{
		ID:        result.ID,
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
	}

	writeJSON(w, http.StatusCreated, response)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	result, err := h.service.Login(r.Context(), req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	h.setRefreshTokenCookie(w, result.RefreshToken)

	response := LoginResponse{
		UserID:      result.UserID,
		AccessToken: result.AccessToken,
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(h.refreshCookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			writeError(
				w,
				http.StatusUnauthorized,
				"refresh token is required",
			)
			return
		}

		writeError(
			w,
			http.StatusBadRequest,
			"invalid refresh token cookie",
		)
		return
	}

	result, err := h.service.Refresh(
		r.Context(),
		cookie.Value,
	)
	if err != nil {
		h.handleError(w, err)
		return
	}

	h.setRefreshTokenCookie(w, result.RefreshToken)

	response := RefreshResponse{
		UserID:      result.UserID,
		AccessToken: result.AccessToken,
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) setRefreshTokenCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     h.refreshCookieName,
		Value:    token,
		Path:     "/auth",
		MaxAge:   h.refreshCookieMaxAge,
		HttpOnly: true,
		Secure:   h.secureCookie,
		SameSite: http.SameSiteLaxMode,
	})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(h.refreshCookieName)

	if err == nil {
		if err := h.service.Logout(
			r.Context(),
			cookie.Value,
		); err != nil {
			h.handleError(w, err)
			return
		}
	} else if !errors.Is(err, http.ErrNoCookie) {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid refresh token cookie",
		)
		return
	}

	h.clearRefreshTokenCookie(w)

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Revoke(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(h.refreshCookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		writeError(
			w,
			http.StatusBadRequest,
			"invalid refresh token cookie",
		)
		return
	}

	if err := h.service.Revoke(
		r.Context(),
		cookie.Value,
	); err != nil {
		h.handleError(w, err)
		return
	}

	h.clearRefreshTokenCookie(w)

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) clearRefreshTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     h.refreshCookieName,
		Value:    "",
		Path:     "/auth",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   h.secureCookie,
		SameSite: http.SameSiteLaxMode,
	})
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
	slog.Error("request failed", "error", err)

	switch {
	case errors.Is(err, ErrEmailRequired),
		errors.Is(err, ErrInvalidEmail),
		errors.Is(err, ErrPasswordRequired),
		errors.Is(err, ErrPasswordTooShort),
		errors.Is(err, ErrPasswordTooLong),
		errors.Is(err, ErrNameRequired),
		errors.Is(err, ErrNameTooLong):
		writeError(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, ErrEmailAlreadyExists):
		writeError(w, http.StatusConflict, err.Error())

	case errors.Is(err, ErrInvalidCredentials):
		writeError(w, http.StatusUnauthorized, "invalid credentials")

	case errors.Is(err, ErrInvalidRefreshToken),
		errors.Is(err, ErrRefreshTokenExpired),
		errors.Is(err, ErrRefreshTokenRevoked):
		writeError(w, http.StatusUnauthorized, "invalid refresh token")

	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}
