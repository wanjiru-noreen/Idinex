package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config stores the application's configuration values.
type Config struct {
	// Port the HTTP server listens on.
	Port string

	// DatabaseURL is used to connect to PostgreSQL.
	DatabaseURL string
}

// Load loads the application's configuration.
func Load() (*Config, error) {

	// Load variables from the .env file if it exists.
	// In production, environment variables are usually provided directly.
	err := godotenv.Load()
	if err != nil {
		// Ignore the error because the .env file is optional.
		// Docker and production environments may not use one.
	}

	// Create a configuration object using environment variables.
	config := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}

	// Return the loaded configuration.
	return config, nil
}

// getEnv retrieves the value of an environment variable.
// If the variable is not set, it returns the fallback value.
func getEnv(key string, fallback string) string {
	// Read the environment variable.
	value := os.Getenv(key)

	// Return the fallback if the variable is empty.
	if value == "" {
		return fallback
	}

	// Return the environment variable's value.
	return value
}
