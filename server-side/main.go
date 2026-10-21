package main

import (
	"log"

	"github.com/kishanravi887321/pumkin/server-side/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
