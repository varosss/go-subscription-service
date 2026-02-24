package config

import "os"

type Database struct {
	DSN string
}

type HttpServer struct {
	Port string
}

type Config struct {
	Database   Database
	HttpServer HttpServer
}

func Load() (*Config, error) {
	cfg := &Config{
		Database: Database{
			DSN: os.Getenv("POSTGRES_URL"),
		},
		HttpServer: HttpServer{
			Port: getEnv("HTTP_SERVER_PORT", "80"),
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
