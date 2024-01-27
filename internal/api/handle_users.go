package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/avearmin/gorage-sale/internal/database"
	"github.com/google/uuid"
)

func (cfg config) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cfg.postUsers(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func (cfg config) postUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := decoder.Decode(&parameters); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if parameters.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid name")
		return
	}
	if ok := verifyEmail(parameters.Email); !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid Email format")
		return
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
		return
	}

	respondWithJSON(w, http.StatusOK, dbUserToJSONUser(user))
}

func verifyEmail(email string) bool {
	emailFields := strings.Split(email, "@")
	if len(emailFields) != 2 {
		return false
	}

	username := emailFields[0]
	domain := emailFields[1]
	if username == "" || domain == "" {
		return false
	}

	domainFields := strings.Split(domain, ".")
	if len(domainFields) < 2 {
		return false
	}

	mailServer := domainFields[0]
	topLevelDomain := domainFields[1]
	if mailServer == "" || topLevelDomain == "" {
		return false
	}

	return true
}
