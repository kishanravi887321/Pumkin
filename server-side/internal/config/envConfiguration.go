package config

import (
	"errors"
	"log"

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
