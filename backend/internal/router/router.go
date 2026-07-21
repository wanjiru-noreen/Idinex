package router

import (
	"encoding/json"
	"net/http"
)

// New creates the application router.
func New() *http.ServeMux {
	router := http.NewServeMux()

	// Root endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Idinex is now in development"))
	})

	// Health endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	return router
}
