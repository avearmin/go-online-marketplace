package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/avearmin/gorage-sale/internal/database"
	"github.com/google/uuid"
)

func (cfg apiConfig) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cfg.postUsers(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func (cfg apiConfig) postUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := decoder.Decode(&parameters); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parameters.Name,
		Email:     parameters.Email,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating user in DB")
	}

	respondWithJSON(w, http.StatusOK, dbUserToJSONUser(user))
}
