package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"idinex-go/config"
)

func main() {
	cfg := config.Load()

	if err := cfg.WaitForDatabase(30 * time.Second); err != nil {
		log.Printf("database connection unavailable: %v", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":      "ok",
			"service":     "backend",
			"environment": cfg.Environment,
		})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Idinex backend is running"))
	})

	address := cfg.HTTPAddress()
	log.Printf("starting backend on %s", address)
	log.Fatal(http.ListenAndServe(address, mux))
}
