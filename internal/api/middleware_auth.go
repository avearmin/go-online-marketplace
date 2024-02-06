package api

import (
	"errors"
	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/avearmin/gorage-sale/internal/database"
	"net/http"
)

type authedUser struct {
	Error error
	User  database.User
}

type authedHandler func(http.ResponseWriter, *http.Request, authedUser)

func (cfg config) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := cfg.authenticateRequest(w, r)
		handler(w, r, user)
	}
}

func (cfg config) authenticateRequest(w http.ResponseWriter, r *http.Request) authedUser {
	accessCookie, err := r.Cookie("gorage-sale-access-token")
	if err != nil {
		return authedUser{Error: err}
	}

	accessToken := accessCookie.Value
	id, err := auth.ValidateAccessToken(accessToken, cfg.JwtSecret)
	if (err != nil && errors.Is(err, auth.ErrTokenExpired)) || isCookieExpired(accessCookie) {
		id, err = refreshAccessToken(w, r, cfg.JwtSecret)
		if err != nil {
			return authedUser{Error: err}
		}
	} else {
		return authedUser{Error: err}
	}

	user, err := cfg.DB.GetUserById(r.Context(), id)
	if err != nil {
		return authedUser{Error: err}
	}

	return authedUser{User: user}
}
