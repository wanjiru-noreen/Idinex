package config

import (
	"testing"
)

func TestLoadUsesEnvironmentOverrides(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_URL", "postgres://postgres:postgres@postgres:5432/idinex?sslmode=disable")
	t.Setenv("JWT_SECRET", "test-secret")
	t.Setenv("APP_ENV", "test")
	t.Setenv("SMTP_PORT", "2525")

	cfg := Load()

	if cfg.Port != "9090" {
		t.Fatalf("expected port override to be applied, got %q", cfg.Port)
	}
	if cfg.DatabaseURL != "postgres://postgres:postgres@postgres:5432/idinex?sslmode=disable" {
		t.Fatalf("expected database URL override to be applied, got %q", cfg.DatabaseURL)
	}
	if cfg.JwtSecret != "test-secret" {
		t.Fatalf("expected JWT_SECRET override to be applied, got %q", cfg.JwtSecret)
	}
	if cfg.SMTPPort != 2525 {
		t.Fatalf("expected SMTP_PORT override to be applied, got %d", cfg.SMTPPort)
	}
}

func TestDatabaseHostAndPortUsesDefaultPortWhenMissing(t *testing.T) {
	cfg := Config{DatabaseURL: "postgres://postgres:postgres@postgres/idinex?sslmode=disable"}
	host, port, err := cfg.DatabaseHostAndPort()
	if err != nil {
		t.Fatalf("expected database URL to parse, got error: %v", err)
	}
	if host != "postgres" {
		t.Fatalf("expected parsed host to be postgres, got %q", host)
	}
	if port != "5432" {
		t.Fatalf("expected default postgres port, got %q", port)
	}
}
