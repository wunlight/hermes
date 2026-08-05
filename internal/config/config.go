package config

type Config struct {
	App  AppConfig
	HTTP HTTPConfig
	DB   DatabaseConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type HTTPConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}
