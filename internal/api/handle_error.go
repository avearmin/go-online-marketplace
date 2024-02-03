package api

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func HandleError(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getErrorPage(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getErrorPage(w http.ResponseWriter, r *http.Request) {
	errorTmplPath := "./web/templates/error.gohtml"
	defaultErrorTmplPath := "./web/templates/default-error.html"

	code := r.URL.Query().Get("code")
	message := r.URL.Query().Get("message")

	unescapedMessage, err := url.QueryUnescape(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, defaultErrorTmplPath)
		return
	}

	var data struct {
		Code    string
		Message string
	}
	data.Code = code
	data.Message = unescapedMessage

	codeInt, err := strconv.Atoi(code)
	if err != nil { // The only reason why this should be an error is if a user is writing a non-status code in the url
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, defaultErrorTmplPath)
		return
	}

	if err := respondWithHTMLFromFile(w, codeInt, errorTmplPath, data); err != nil {
		log.Printf("Server failed to respond with error page: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

}
