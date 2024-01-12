package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/avearmin/gorage-sale/internal/database"
	"github.com/google/uuid"
)

func (cfg apiConfig) handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cfg.middlewareAuth(cfg.postItems).ServeHTTP(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func (cfg apiConfig) postItems(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var parameters struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int32  `json:"price"`
	}
	if err := decoder.Decode(&parameters); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if ok := validateItemName(parameters.Name); !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid name")
		return
	}
	if ok := validateItemDescription(parameters.Description); !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid description")
		return
	}
	if parameters.Price < 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid price")
		return
	}

	item, err := cfg.DB.CreateItem(r.Context(), database.CreateItemParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        parameters.Name,
		Description: parameters.Description,
		Price:       parameters.Price,
		Sold:        false,
		SellerID:    id,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, dbItemToJSONItem(item))
}

func validateItemName(name string) bool {
	return name != "" && len(name) <= 72
}

func validateItemDescription(description string) bool {
	return description != "" && len(description) <= 720
}
