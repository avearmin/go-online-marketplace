package main

import (
	"encoding/json"
	"net/http"
	//	"github.com/avearmin/gorage-sale/internal/auth"
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
		Email string `json:"email"`
	}
	if err := decoder.Decode(&parameters); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid JSON format")
		return
	}

	// googleJWT, err := getGoogleJWT(parameters.Email)
	// if err != nil {
	//	respondWithError(w, http.StatusUnauthorized, "failed to login to google")
	//	return
	// }

	// pem, err := auth.GetGooglePublicKey(fmt.Sprintf("%s", googleJWT.Header["kid"]))
	// if err != nil {
	//	if err.Error() == "key not found" {
	//		respondWithError(w, http.StatusNotFound, "public key not found for provided key identifier")
	//		return
	//	}
	//	respondWithError(w, http.StatusInternalServiceError, "error getting google public key")
	//	return
	// }
	// claims, err := auth.ValidateGoogleJWT(googleJWT, pem, cfg.ClientID)
	// if err != nil {
	//	respondWithError(w, http.StatusForbidden, "invalid google auth")
	//	return
	// }

	// id, err := cfg.DB.GetUserIDByEmail(parameters.Email)
	// if err != nil {
	//	if err == sql.ErrNoRows {
	//		user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
	//			ID:        uuid.New(),
	//			CreatedAt: time.Now().UTC(),
	//			UpdatedAt: time.Now().UTC(),
	//			Name:      claims.Name,
	//			Email:     parameters.Email,
	//		})
	//		if err != nil {
	//			respondWithError(w, http.InternalServiceError, "error creating user in database")
	//			return
	//		}
	//		id = user.ID
	//	} else {
	//		respondWithError(w, http.InternalServiceError, "error creating user in database")
	//		return
	//	}
	// }

	// accessToken, err := auth.CreateAccessToken(id, cfg.JwtSecret)
	// if err != nil {
	//	respondWithError(w, http.InternalServerError, "error creating access token")
	//	return
	//}
	// refreshToken, err := auth.CreateRefreshToken(id, cfg.JwtSecret)
	// if err != nil {
	//	respondWithError(w, http.InternalServerError, "error creating refresh token")
	//	return
	// }

	// var payload struct {
	//	AccessToken string `json:"access_token"`
	//	RefreshToken string `json:"refresh_token"`
	// }
	// payload.AccessToken = accessToken
	// payload.RefreshToken = refreshToken
	//
	// respondWithJSON(w, http.StatusOK, payload)
}
