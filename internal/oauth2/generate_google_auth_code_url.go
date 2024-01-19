package oauth2

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GenerateGoogleAuthCodeURL(clientID, clientSecret, redirectURL, state string) string {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	url := conf.AuthCodeURL(state)
	return url
}
