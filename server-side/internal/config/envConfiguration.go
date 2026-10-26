package config

import (
	"errors"
	"log"
	"time"

	// "log"
	"os"
)

//  get the data form the .env file
// and return from this file

func Dburl() (string, error) {

	url := os.Getenv("MONGO__DB__URL")

	if url == "" {

		log.Println("eror occured while readign the databse url")
		return "", errors.New("MONGO__DB__URL is not configured")
	}

	// fmt.Println(url)

	return url, nil

}

func NeonDburl() (string, error) {

	url := os.Getenv("NEON__DB__URL")

	if url == "" {
		log.Println("error occurred while reading the Neon database URL")
		return "", errors.New("NEON__DB__URL is not configured")
	}

	return url, nil
}

func GetPort() (string, error) {

	port := os.Getenv("PORT")

	if port == "" {
		return "", errors.New("Port is not configured")
	}

	return port, nil

}

func Getenvironment() (string, error) {
	env := os.Getenv("ENV")

	if env == "" {
		return "", errors.New("Env is not configured")
	}
	return env, nil
}

func GetAccessTokenSecret() (string, error) {
	secret := os.Getenv("ACCESS__TOKEN__SECRET")
	if secret == "" {
		return "", errors.New("ACCESS__TOKEN__SECRET is not configured")
	}
	return secret, nil
}

func GetRefreshTokenSecret() (string, error) {
	secret := os.Getenv("REFRESH__TOKEN__SECRET")
	if secret == "" {
		return "", errors.New("REFRESH__TOKEN__SECRET is not configured")
	}
	return secret, nil
}

func GetAccessTokenExpiry() (time.Duration, error) {
	expiry := os.Getenv("ACCESS__TOKEN__EXPIRY")
	if expiry == "" {
		return 0, errors.New("ACCESS__TOKEN__EXPIRY is not configured")
	}
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		return 0, errors.New("ACCESS__TOKEN__EXPIRY is not a valid duration")
	}
	return duration, nil
}

func GetRefreshTokenExpiry() (time.Duration, error) {
	expiry := os.Getenv("REFRESH__TOKEN__EXPIRY")
	if expiry == "" {
		return 0, errors.New("REFRESH__TOKEN__EXPIRY is not configured")
	}
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		return 0, errors.New("REFRESH__TOKEN__EXPIRY is not a valid duration")
	}
	return duration, nil
}

func GetEncryptionKey() (string, error) {
	key := os.Getenv("MASTER__KEY")
	if key == "" {
		return "", errors.New("MASTER__KEY is not configured")
	}

	return key, nil
}
