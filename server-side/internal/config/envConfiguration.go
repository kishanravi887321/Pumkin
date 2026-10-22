package config

import (
	"log"
	// "log"
	"os"
)

//  get the data form the .env file
// and return from this file

func Dburl() string {

	url := os.Getenv("MONGO__DB__URL")

	if url == "" {
		log.Fatal("eror occured while readign the databse url")
	}

	// fmt.Println(url)

	return url

}

func GetPort() string {

	return os.Getenv("PORT")
}

func Getenvironment() string {
	return os.Getenv("ENV")
}
