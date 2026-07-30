package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"idinex-go/config"
	"idinex-go/internal/database"
	"idinex-go/internal/router"
)

func main() {
	// Load application configuration.
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the database to become available.
	if err := cfg.WaitForDatabase(30 * time.Second); err != nil {
		log.Printf("database connection unavailable: %v", err)
		os.Exit(1)
	}

	// Connect to PostgreSQL.
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create application router.
	appRouter := router.New()

	// Health endpoint.
	appRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":      "ok",
			"service":     "backend",
			"environment": cfg.Environment,
		})
	})

	// Root endpoint (only if router.New() doesn't already register one).
	appRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Idinex backend is running"))
	})

	address := cfg.HTTPAddress()
	log.Printf("starting backend on %s", address)

	if err := http.ListenAndServe(address, appRouter); err != nil {
		log.Fatal(err)
	}
}