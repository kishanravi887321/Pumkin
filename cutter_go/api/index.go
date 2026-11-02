package handler

import (
	"net/http"

	cuttergo "github.com/kishanravi887321/pumkin/cutter_go"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	cuttergo.Handler(response, request)
}
