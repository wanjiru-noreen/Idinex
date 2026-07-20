package config

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

// DatabaseHostAndPort extracts the host and port from the configured DATABASE_URL.
func (c Config) DatabaseHostAndPort() (string, string, error) {
	parsed, err := url.Parse(c.DatabaseURL)
	if err != nil {
		return "", "", fmt.Errorf("parse database URL: %w", err)
	}

	host := parsed.Hostname()
	if strings.TrimSpace(host) == "" {
		return "", "", fmt.Errorf("database host is empty")
	}

	port := parsed.Port()
	if port == "" {
		port = "5432"
	}

	return host, port, nil
}

// WaitForDatabase repeatedly attempts a TCP connection to the configured database host.
func (c Config) WaitForDatabase(timeout time.Duration) error {
	host, port, err := c.DatabaseHostAndPort()
	if err != nil {
		return err
	}

	deadline := time.Now().Add(timeout)
	var lastErr error
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 2*time.Second)
		if err == nil {
			_ = conn.Close()
			return nil
		}
		lastErr = err
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("database not reachable at %s:%s: %w", host, port, lastErr)
}
