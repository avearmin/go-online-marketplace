package oauth2

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

type GoogleUserData struct {
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
}

func GetUserDataFromGoogle(clientID, clientSecret, redirectURL, code string, ctx context.Context) (GoogleUserData, error) {
	oauthGoogleUrlAPI := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return GoogleUserData{}, err
	}

	response, err := http.NewRequestWithContext(ctx, "GET", oauthGoogleUrlAPI+token.AccessToken, nil)
	if err != nil {
		return GoogleUserData{}, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var userData GoogleUserData
	if err := decoder.Decode(&userData); err != nil {
		return GoogleUserData{}, err
	}

	return userData, nil
}
