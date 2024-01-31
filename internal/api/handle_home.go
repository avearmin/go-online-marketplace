package api

import (
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getHomePage(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "")
	}
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/index.html")
}
