package main

import (
	"errors"
	"log"
	"net/http"

	cuttergo "github.com/kishanravi887321/pumkin/cutter_go"
)

func main() {
	if err := cuttergo.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
