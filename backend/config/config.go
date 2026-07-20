package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config stores application configuration.
type Config struct {
	Port        string
	DatabaseURL string
}

// Load loads configuration values.
func Load() (*Config, error) {

	// Loads variables from .env file if it exists.
	// In production, environment variables are usually provided directly.
	err := godotenv.Load()
	if err != nil {
		// Ignore error because .env is optional.
		// Docker and production environments may not use a .env file.
	}

	config := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}

	return config, nil
}

// getEnv retrieves an environment variable.
// If the variable does not exist, it returns a fallback value.
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
