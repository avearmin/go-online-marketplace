package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/avearmin/gorage-sale/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg config) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := readAccessToken(r)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := auth.ValidateAccessToken(accessToken, cfg.JwtSecret)
		if err != nil {
			if err == auth.ErrTokenExpired {
				respondWithError(w, http.StatusUnauthorized, err.Error())
				return
			}
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := cfg.DB.GetUserById(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				respondWithError(w, http.StatusNotFound, err.Error())
				return
			}
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, user)
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
