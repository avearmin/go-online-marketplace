package main

import (
	"github.com/avearmin/gorage-sale/internal/oauth2"
	"net/http"
)

func (cfg apiConfig) handleOAuthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cfg.getOAuthGoogleLogin(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func (cfg apiConfig) getOAuthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState, err := cfg.StateStore.GenerateState()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error generating state")
		return
	}
	url := oauth2.GenerateGoogleAuthCodeURL(cfg.ClientID, cfg.ClientSecret, cfg.OAuthRedirectURL, oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
