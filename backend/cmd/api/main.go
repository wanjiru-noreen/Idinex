package main

import (
	"fmt"
	"log"
	"net/http"

	"idinex-go/config"
	"idinex-go/internal/database"
	"idinex-go/internal/router"
)

func main() {
	// Load application configuration from the .env file.
	cfg, err := config.Load()

	// Stop the application if configuration fails to load.
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the PostgreSQL database.
	db, err := database.Connect(cfg.DatabaseURL)

	// Stop the application if the database connection fails.
	if err != nil {
		log.Fatal(err)
	}

	// Close the database connection when the application exits.
	defer db.Close()

	// Create and register all application routes.
	appRouter := router.New()

	// Display the port the server is running on.
	fmt.Println("Server running on port", cfg.Port)

	// Start the HTTP server.
	err = http.ListenAndServe(":"+cfg.Port, appRouter)

	// Stop the application if the server fails to start.
	if err != nil {
		log.Fatal(err)
	}
}
