package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// Config contains the environment-driven settings used by the API.
type Config struct {
	Port         string
	DatabaseURL  string
	JwtSecret    string
	Environment  string
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

// Load reads configuration from the environment.
// It attempts to load a .env file first, but does not fail if one is not present.
func Load() (*Config, error) {
	// Ignore the error because .env is optional.
	_ = godotenv.Load()

	cfg := &Config{
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://postgres:postgres@postgres:5432/idinex?sslmode=disable"),
		JwtSecret:    getEnv("JWT_SECRET", "dev-secret"),
		Environment:  getEnv("APP_ENV", getEnv("ENVIRONMENT", "development")),
		SMTPHost:     getEnv("SMTP_HOST", "localhost"),
		SMTPPort:     getEnvInt("SMTP_PORT", 1025),
		SMTPUser:     getEnv("SMTP_USER", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
	}
	return fallback
}

// HTTPAddress normalizes the configured port into an address suitable for net/http.
func (c *Config) HTTPAddress() string {
	if strings.HasPrefix(c.Port, ":") {
		return c.Port
	}
	return fmt.Sprintf(":%s", c.Port)
}


