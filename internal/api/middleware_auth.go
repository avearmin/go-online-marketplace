package api

import (
	"github.com/avearmin/gorage-sale/internal/auth"
	"github.com/avearmin/gorage-sale/internal/database"
	"net/http"
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
	accessCookie, err := r.Cookie("gorage-sale-access-token")
	if err != nil {
		return "", err
	}

	return accessCookie.Value, nil
}
