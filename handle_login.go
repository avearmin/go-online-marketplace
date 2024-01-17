package main

import (
	"encoding/json"
	"net/http"
)

func (cfg apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cfg.getLogin(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func (cfg apiConfig) getLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var parameters struct {
		GoogleJWT string `json:"google_jwt"`	
	}
	if err := decoder.Decode(&parameters); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}



}
