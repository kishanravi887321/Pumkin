package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kishanravi887321/pumkin/server-side/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found; using system environment variables")
	} else {
		log.Println(".env is loaded ......")
	}

	err = cmd.Run()
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
