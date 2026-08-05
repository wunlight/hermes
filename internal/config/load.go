package config

import "fmt"

func Load() (*Config, error) {
	httpCfg, err := loadHTTP()
	if err != nil {
		return nil, err
	}

	dbCfg, err := loadDatabase()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App:  loadApp(),
		HTTP: httpCfg,
		DB:   dbCfg,
	}

	return cfg, nil
}

func loadApp() AppConfig {
	return AppConfig{
		Name: stringEnv("APP_NAME", "hermes"),
		Env:  stringEnv("APP_ENV", "development"),
	}
}

func loadHTTP() (HTTPConfig, error) {
	port, err := intEnv("HTTP_PORT", 8080)
	if err != nil {
		return HTTPConfig{}, fmt.Errorf("load http config: %w", err)
	}
	return HTTPConfig{Port: port}, nil
}

func loadDatabase() (DatabaseConfig, error) {
	host, err := requiredString("DB_HOST")
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("load database config: %w", err)
	}

	port, err := requiredInt("DB_PORT")
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("load database config: %w", err)
	}

	user, err := requiredString("DB_USER")
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("load database config: %w", err)
	}

	password, err := requiredString("DB_PASSWORD")
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("load database config: %w", err)
	}

	name, err := requiredString("DB_NAME")
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("load database config: %w", err)
	}

	return DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
		SSLMode:  stringEnv("DB_SSLMODE", "disable"),
	}, nil
}
