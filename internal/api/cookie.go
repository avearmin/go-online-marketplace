package api

import (
	"errors"
	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var (
	errorNoRefreshToken = errors.New("no cookie with refresh token found")
)

func setAccessCookie(w http.ResponseWriter, token string) {
	accessCookie := http.Cookie{
		Name:    "gorage-sale-access-token",
		Value:   token,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	}

	http.SetCookie(w, &accessCookie)
}

func setRefreshCookie(w http.ResponseWriter, token string) {
	refreshCookie := http.Cookie{
		Name:    "gorage-sale-refresh-token",
		Value:   token,
		Expires: time.Now().Add(1440 * time.Hour),
		Path:    "/",
	}

	http.SetCookie(w, &refreshCookie)
}

func isCookieExpired(c *http.Cookie) bool {
	if c == nil {
		return true
	}
	return time.Now().After(c.Expires)
}

func refreshAccessToken(w http.ResponseWriter, r *http.Request, jwtSecret string) (uuid.UUID, error) {
	refreshCookie, err := r.Cookie("gorage-sale-refresh-token")
	if err != nil {
		return uuid.Nil, errorNoRefreshToken
	}

	refreshToken := refreshCookie.Value

	id, err := auth.ValidateRefreshToken(refreshToken, jwtSecret)
	if err != nil {
		return uuid.Nil, err
	}

	accessToken, err := auth.CreateAccessToken(id, jwtSecret)
	if err != nil {
		return uuid.Nil, err
	}

	setAccessCookie(w, accessToken)
	return id, nil
}
