package config

import "os"

type Config struct {
	Port        string
	Environment string
	DatabaseURL string
}

func Load() Config {
	return Config{
		Port:        GetPort(),
		Environment: Getenvironment(),
		DatabaseURL: Dburl(),
	}
}

func valueOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
