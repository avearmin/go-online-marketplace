package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/avearmin/gorage-sale/internal/database"
	"github.com/avearmin/gorage-sale/internal/oauth2"
	"github.com/google/uuid"
)

func (cfg config) HandleOAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cfg.getOAuthGoogleCallback(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (cfg config) getOAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState := r.FormValue("state")

	if valid := cfg.StateStore.ValidateState(oauthState); !valid {
		redirectToErrorPage(w, r, http.StatusUnauthorized, "Access Denied")
		return
	}
	cfg.StateStore.DeleteState(oauthState)

	code := r.FormValue("code")
	data, err := oauth2.GetUserDataFromGoogle(cfg.ClientID, cfg.ClientSecret, cfg.OAuthRedirectURL, code, r.Context())
	if err != nil {
		logInternalErrorAndRedirectToErrorPage(w, r, err)
		return
	}

	id, err := cfg.DB.GetUserIDByEmail(r.Context(), data.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // If the user does not exist, then we create one
			user, err := cfg.createUser(r.Context(), data.Name, data.Email)
			if err != nil {
				logInternalErrorAndRedirectToErrorPage(w, r, err)
				return
			}
			id = user.ID
		} else {
			logInternalErrorAndRedirectToErrorPage(w, r, err)
			return
		}
	}

	accessToken, err := auth.CreateAccessToken(id, cfg.JwtSecret)
	if err != nil {
		logInternalErrorAndRedirectToErrorPage(w, r, err)
		return
	}
	refreshToken, err := auth.CreateRefreshToken(id, cfg.JwtSecret)
	if err != nil {
		logInternalErrorAndRedirectToErrorPage(w, r, err)
		return
	}

	accessCookie := http.Cookie{
		Name:    "gorage-sale-access-token",
		Value:   accessToken,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	}
	refreshCookie := http.Cookie{
		Name:    "gorage-sale-refresh-token",
		Value:   refreshToken,
		Expires: time.Now().Add(1440 * time.Hour),
		Path:    "/",
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (cfg config) createUser(ctx context.Context, name, email string) (database.User, error) {
	user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Email:     email,
	})
	return user, err
}
