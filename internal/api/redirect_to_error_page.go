package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	internalErrMsgForUser = `Seems like we're the ones with the problem. We're looking into it.`
)

func redirectToErrorPage(w http.ResponseWriter, r *http.Request, code int, message string) {
	escapedMessage := url.QueryEscape(message)

	url := fmt.Sprintf("/error?code=%d&message=%s", code, escapedMessage)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// a wrapper so I don't have to make a redirectToErrorPage function with 5 parameters just incase
// we have a 500 code.
func logInternalErrorAndRedirectToErrorPage(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Server responding with code 500: %s", err.Error())
	redirectToErrorPage(w, r, http.StatusInternalServerError, internalErrMsgForUser)
}
