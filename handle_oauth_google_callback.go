package main

import (
	"github.com/avearmin/gorage-sale/internal/oauth2"
	"net/http"
)

func (cfg apiConfig) handleOAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cfg.getOAuthGoogleCallback(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func (cfg apiConfig) getOAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState := r.FormValue("state")

	if valid := cfg.StateStore.ValidateState(oauthState); !valid {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	cfg.StateStore.DeleteState(oauthState)

	code := r.FormValue("code")
	data, err := oauth2.GetUserDataFromGoogle(cfg.ClientID, cfg.ClientSecret, cfg.OAuthRedirectURL, code, r.Context())
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
}
