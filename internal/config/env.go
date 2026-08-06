package config

import (
	"fmt"
	"os"
	"strconv"
)

func optionalString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func requiredString(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return "", fmt.Errorf("%q required", key)
	}
	return value, nil
}

func optionalInt(key string, fallback int) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback, nil
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%q must be an integer: %w", key, err)
	}

	return i, nil
}

func requiredInt(key string) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return 0, fmt.Errorf("%q required", key)
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%q must be an integer: %w", key, err)
	}

	return i, nil
}
