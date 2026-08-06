package config

import "fmt"

func Load() (*Config, error) {
	cfg := &Config{}
	var err error

	cfg.App = loadApp()

	if cfg.HTTP, err = loadHTTP(); err != nil {
		return nil, err
	}

	if cfg.DB, err = loadDatabase(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadApp() AppConfig {
	return AppConfig{
		Name: optionalString("APP_NAME", "hermes"),
		Env:  optionalString("APP_ENV", "development"),
	}
}

func loadHTTP() (HTTPConfig, error) {
	cfg := HTTPConfig{}
	var err error

	if cfg.Port, err = optionalInt("HTTP_PORT", 8080); err != nil {
		return cfg, fmt.Errorf("http config: %w", err)
	}

	return cfg, nil
}

func loadDatabase() (DatabaseConfig, error) {
	cfg := DatabaseConfig{}
	var err error

	if cfg.Host, err = requiredString("DB_HOST"); err != nil {
		return cfg, fmt.Errorf("database config: %w", err)
	}

	if cfg.Port, err = requiredInt("DB_PORT"); err != nil {
		return cfg, fmt.Errorf("database config: %w", err)
	}

	if cfg.User, err = requiredString("DB_USER"); err != nil {
		return cfg, fmt.Errorf("database config: %w", err)
	}

	if cfg.Password, err = requiredString("DB_PASSWORD"); err != nil {
		return cfg, fmt.Errorf("database config: %w", err)
	}

	if cfg.Name, err = requiredString("DB_NAME"); err != nil {
		return cfg, fmt.Errorf("database config: %w", err)
	}

	cfg.SSLMode = optionalString("DB_SSLMODE", "disable")

	return cfg, nil
}
