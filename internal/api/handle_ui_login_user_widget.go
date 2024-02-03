package api

import (
	"net/http"
)

func (cfg config) HandleUILoginUserWidget(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cfg.middlewareAuth(getLoginUserWidget).ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func getLoginUserWidget(w http.ResponseWriter, r *http.Request, user authedUser) {
	tmplPath := "./web/templates/partials/user-login-widget.gohtml"
	if err := respondWithHTMLFromFile(w, http.StatusOK, tmplPath, user); err != nil {
		redirectToErrorPage(w, r, http.StatusInternalServerError, internalErrMsgForUser)
	}
}
