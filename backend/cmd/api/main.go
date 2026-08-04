package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"idinex-go/config"
	"idinex-go/internal/database"
	"idinex-go/internal/middleware"
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

	// TODO: Pass db into your handlers/services once they are implemented.
	_ = db

	// Create the application router.
	appRouter := router.New()

	handler := middleware.CORS(appRouter)

	// Start the HTTP server.
	address := cfg.HTTPAddress()
	log.Printf("starting backend on %s", address)

	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatal(err)
	}
}
