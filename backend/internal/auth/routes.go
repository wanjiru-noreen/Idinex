package auth

import "net/http"

// RegisterRoutes adds the authentication endpoints to the router.
func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/auth/register", h.Register)
}
