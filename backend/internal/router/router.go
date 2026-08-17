package router

import (
	"encoding/json"
	"net/http"

	"idinex-go/internal/auth"
)

// New creates and configures the application's router.
func New() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Idinex is now in development"))
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"service": "backend",
		})
	})

	auth.RegisterRoutes(router)

	return router
}
