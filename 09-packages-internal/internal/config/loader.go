package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds application-wide configuration.
// This is internal to prevent external packages from depending on config structure.
type Config struct {
	DatabaseURL string
	APIKey      string
	Port        int
	Debug       bool
	MaxRetries  int
}

// Load reads configuration from environment variables.
// This function is only available to packages within this module.
func Load() (*Config, error) {
	cfg := &Config{
		DatabaseURL: getEnvOrDefault("DATABASE_URL", "localhost:5432"),
		APIKey:      os.Getenv("API_KEY"),
		Debug:       os.Getenv("DEBUG") == "true",
		MaxRetries:  3, // default value
	}

	// Parse port from environment
	if portStr := os.Getenv("PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT value %q: %w", portStr, err)
		}
		cfg.Port = port
	} else {
		cfg.Port = 8080 // default port
	}

	// Parse max retries if provided
	if retriesStr := os.Getenv("MAX_RETRIES"); retriesStr != "" {
		retries, err := strconv.Atoi(retriesStr)
		if err != nil {
			return nil, fmt.Errorf("invalid MAX_RETRIES value %q: %w", retriesStr, err)
		}
		if retries < 1 {
			return nil, fmt.Errorf("MAX_RETRIES must be at least 1, got %d", retries)
		}
		cfg.MaxRetries = retries
	}

	// Validate required fields
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable is required")
	}

	return cfg, nil
}

// getEnvOrDefault returns the environment variable value or a default if not set.
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// String returns a string representation of the config (without sensitive data).
func (c *Config) String() string {
	return fmt.Sprintf("Config{DatabaseURL: %s, Port: %d, Debug: %t, MaxRetries: %d}",
		c.DatabaseURL, c.Port, c.Debug, c.MaxRetries)
}

// IsProduction returns true if the application is running in production mode.
func (c *Config) IsProduction() bool {
	return !c.Debug
}
