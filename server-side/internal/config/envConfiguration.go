package config

import (
	"fmt"
	"os"
	"time"
)

func GetConfigValue(
	environment string,
	envKey string,
	secretName string,
	version string,
) (string, error) {
	switch environment {
	case "development":
		value := os.Getenv(envKey)
		if value == "" {
			return "", fmt.Errorf("%s is not configured", envKey)
		}
		return value, nil

	case "production":
		return GetSecret(secretName, version)

	default:
		return "", fmt.Errorf(
			"unsupported environment: %q",
			environment,
		)
	}
}

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("PORT is not configured")
	}
	return port, nil
}

func GetAccessTokenSecret() (string, error) {
	return GetConfigValue(os.Getenv("ENV"), "ACCESS__TOKEN__SECRET", "ACCESS__TOKEN__SECRET", "latest")
}

func GetPrivateServerURL() (string, error) {
	return GetConfigValue(os.Getenv("ENV"), "PRIVATE__SERVER__URL", "PRIVATE__SERVER__URL", "latest")
}

func GetAccessTokenExpiry() (time.Duration, error) {
	value, err := GetConfigValue(os.Getenv("ENV"), "ACCESS__TOKEN__EXPIRY", "ACCESS__TOKEN__EXPIRY", "latest")
	if err != nil {
		return 0, err
	}
	return time.ParseDuration(value)
}
