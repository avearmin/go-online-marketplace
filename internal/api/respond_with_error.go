package api

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code/100 == 5 {
		log.Printf("Responding with status code %d: %s", code, msg)
	}

	var errorResponse struct {
		Error string `json:"error"`
	}
	errorResponse.Error = msg

	respondWithJSON(w, code, errorResponse)
}
