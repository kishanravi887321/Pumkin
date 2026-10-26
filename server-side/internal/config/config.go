package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Port                   string
	Environment            string
	DatabaseURL            string
	NeonDBURL              string
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiry      time.Duration
	RefreshTokenExpiry     time.Duration
	EncryptionKey          string
	PrivateServerURL       string
	PrivateServerAuthToken string
	GoogleClientID         string
	GoogleClientSecret     string
}

func Load() (Config, error) {
	environment := os.Getenv("ENV")
	if environment == "" {
		return Config{}, fmt.Errorf("ENV is not configured")
	}
	if environment == "dev" || environment == "development" {
		environment = "development"
	}
	if environment == "prod" || environment == "production" {
		environment = "production"
	}

	port, err := GetPort()
	if err != nil {
		return Config{}, err
	}
	googleClientID, err := GetConfigValue(
		environment,
		"GOOGLE__CLIENT__ID",
		"GOOGLE__CLIENT__ID",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	googleClientSecret, err := GetConfigValue(
		environment,
		"GOOGLE__CLIENT__SECRET",
		"GOOGLE__CLIENT__SECRET",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	databaseURL, err := GetConfigValue(
		environment,
		"MONGO__DB__URL",
		"MONGO__DB__URL",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	neonDBURL, err := GetConfigValue(
		environment,
		"NEON__DB__URL",
		"NEON__DB__URL",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	accessTokenSecret, err := GetConfigValue(
		environment,
		"ACCESS__TOKEN__SECRET",
		"ACCESS__TOKEN__SECRET",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	refreshTokenSecret, err := GetConfigValue(
		environment,
		"REFRESH__TOKEN__SECRET",
		"REFRESH__TOKEN__SECRET",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	masterKey, err := GetConfigValue(
		environment,
		"MASTER__KEY",
		"MASTER__KEY",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	privateServerURL, err := GetConfigValue(
		environment,
		"PRIVATE__SERVER__URL",
		"PRIVATE__SERVER__URL",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	accessExpiryStr, err := GetConfigValue(
		environment,
		"ACCESS__TOKEN__EXPIRY",
		"ACCESS__TOKEN__EXPIRY",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	accessExpiry, err := time.ParseDuration(accessExpiryStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid ACCESS__TOKEN__EXPIRY: %w", err)
	}

	refreshExpiryStr, err := GetConfigValue(
		environment,
		"REFRESH__TOKEN__EXPIRY",
		"REFRESH__TOKEN__EXPIRY",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	refreshExpiry, err := time.ParseDuration(refreshExpiryStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid REFRESH__TOKEN__EXPIRY: %w", err)
	}

	privateServerAuthToken, err := GetConfigValue(
		environment,
		"PRIVATE_SERVER_AUTH_TOKEN",
		"PRIVATE_SERVER_AUTH_TOKEN",
		"latest",
	)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Port:                   port,
		Environment:            environment,
		DatabaseURL:            databaseURL,
		NeonDBURL:              neonDBURL,
		AccessTokenSecret:      accessTokenSecret,
		RefreshTokenSecret:     refreshTokenSecret,
		AccessTokenExpiry:      accessExpiry,
		RefreshTokenExpiry:     refreshExpiry,
		EncryptionKey:          masterKey,
		PrivateServerURL:       privateServerURL,
		PrivateServerAuthToken: privateServerAuthToken,
		GoogleClientID:         googleClientID,
		GoogleClientSecret:     googleClientSecret,
	}, nil
}
