package config

import "os"

type Config struct {
	Port        string
	Environment string
	DatabaseURL string
}

func Load() Config {
	dburl, err := Dburl()
	if err != nil {
		return Config{}
	}
	port, err := GetPort()

	if err != nil {
		return Config{}
	}

	environment, err := Getenvironment()

	if err != nil {
		return Config{}
	}

	return Config{
		Port:        port,
		Environment: environment,
		DatabaseURL: dburl,
	}
}

func valueOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
