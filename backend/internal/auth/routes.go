package auth

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", RegisterHandler)
}
