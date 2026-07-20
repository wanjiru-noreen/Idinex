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
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	appRouter := router.New()

	fmt.Println("Server running on port", cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, appRouter)

	if err != nil {
		log.Fatal(err)
	}
}
