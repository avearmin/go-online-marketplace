package api

import (
	"log"
	"net/http"
)

func respondWithErrorHTML(w http.ResponseWriter, code int, tmplString string, errMsg string) {
	if code/100 == 5 {
		log.Printf("Responding with code %d: %s", code, errMsg)
	}

	var payload struct {
		Error string
	}
	payload.Error = errMsg

	if err := respondWithHTML(w, code, tmplString, payload); err != nil {
		log.Printf("Failed to send error: %s", err.Error())
	}
}

func respondWithErrorHTMLFromFile(w http.ResponseWriter, code int, tmplPath string, errMsg string) {
	if code/100 == 5 {
		log.Printf("Responding with code %d: %s", code, errMsg)
	}

	var payload struct {
		Error string
	}
	payload.Error = errMsg

	if err := respondWithHTMLFromFile(w, code, tmplPath, payload); err != nil {
		log.Printf("Failed to send error: %s", err.Error())
	}
}
