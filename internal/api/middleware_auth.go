package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/google/uuid"
)

type authedHandler func(http.ResponseWriter, *http.Request, uuid.UUID)

func (cfg config) middlewareAuth(handler authedHandler) http.Handler {
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

		if _, err := cfg.DB.GetUserById(r.Context(), id); err != nil {
			if err == sql.ErrNoRows {
				respondWithError(w, http.StatusNotFound, err.Error())
				return
			}
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, id)
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
