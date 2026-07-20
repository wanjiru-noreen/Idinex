package router

import (
	"net/http"
)

// New creates the application router.
func New() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Idinex API running"))
	})

	return router
}
