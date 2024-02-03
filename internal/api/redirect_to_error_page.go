package api

import (
	"fmt"
	"net/http"
	"net/url"
)

func redirectToErrorPage(w http.ResponseWriter, r *http.Request, code int, message string) {
	escapedMessage := url.QueryEscape(message)

	url := fmt.Sprintf("/error/code=%d&message=%s", code, escapedMessage)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
