package router

import (
	"encoding/json"
	"net/http"
)

// New creates and configures the application's router.
func New() *http.ServeMux {
	// Create a new HTTP request multiplexer.
	router := http.NewServeMux()

	// Register the root endpoint.
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Return the application status message.
		w.Write([]byte("Idinex is now in development"))
	})

	// Register the health check endpoint.
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"service": "backend",
		})
	})

	// Return the configured router.
	return router
}
