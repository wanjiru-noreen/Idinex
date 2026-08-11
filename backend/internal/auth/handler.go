package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Handler receives HTTP requests for the authentication endpoints.
type Handler struct {
	svc *Service
}

// NewHandler creates a Handler backed by the given Service.
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// Register handles POST /auth/register.
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if msg := validateRegister(req); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}

	user, err := h.svc.RegisterUser(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, ErrEmailTaken), errors.Is(err, ErrUsernameTaken):
			writeError(w, http.StatusConflict, err.Error())
		default:
			log.Printf("register: %v", err)
			writeError(w, http.StatusInternalServerError, "registration failed")
		}
		return
	}

	writeJSON(w, http.StatusCreated, RegisterResponse{
		Message: "account created",
		User:    user,
	})
}

// writeJSON sends a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// writeError sends a JSON error response with the given status code.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
