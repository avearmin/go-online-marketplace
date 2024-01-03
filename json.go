package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		log.Printf("Error responding with JSON: %s", err)
	}
}

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
