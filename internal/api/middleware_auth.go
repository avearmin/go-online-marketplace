package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/avearmin/gorage-sale/internal/database"
)

type authedUser struct {
	IsAuthed bool
	User     database.User
}

type authedHandler func(http.ResponseWriter, *http.Request, authedUser)

func (cfg config) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuthed := true // default to true, and we set it to false for any error

		accessToken, err := readAccessToken(r)
		if err != nil {
			isAuthed = false
		}

		id, err := auth.ValidateAccessToken(accessToken, cfg.JwtSecret)
		if err != nil {
			isAuthed = false
		}

		user, err := cfg.DB.GetUserById(r.Context(), id)
		if err != nil {
			isAuthed = false
		}

		authUser := authedUser{
			IsAuthed: isAuthed,
			User:     user,
		}

		handler(w, r, authUser)
	})
}

func readAccessToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return "", errors.New("malformed authorization header")
	}
	if fields[0] != "Bearer" {
		return "", errors.New("bearer not found in header")
	}
	return fields[1], nil
}
