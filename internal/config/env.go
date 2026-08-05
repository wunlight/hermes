package config

import (
	"fmt"
	"os"
	"strconv"
)

func stringEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func requiredString(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return "", fmt.Errorf("required environment variable %q is not set", key)
	}
	return value, nil
}

func intEnv(key string, fallback int) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback, nil
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("environment variable %q must be an integer: %w", key, err)
	}

	return i, nil
}

func requiredInt(key string) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return 0, fmt.Errorf("required environment variable %q is not set", key)
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("environment variable %q must be an integer: %w", key, err)
	}

	return i, nil
}
