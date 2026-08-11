package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"idinex-go/config"
	"idinex-go/internal/auth"
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

	// Create application router. It already registers the root and
	// /health endpoints.
	appRouter := router.New()

	// Register the authentication module.
	authHandler := auth.NewHandler(auth.NewService(auth.NewRepository(db)))
	auth.RegisterRoutes(appRouter, authHandler)

	address := cfg.HTTPAddress()
	log.Printf("starting backend on %s", address)

	if err := http.ListenAndServe(address, appRouter); err != nil {
		log.Fatal(err)
	}
}
