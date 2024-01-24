package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
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
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return GoogleUserData{}, err
	}

	oauthGoogleUrlAPI := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="

	fmt.Println(token.AccessToken)

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return GoogleUserData{}, err
	}

	decoder := json.NewDecoder(response.Body)

	var userData GoogleUserData
	if err := decoder.Decode(&userData); err != nil {
		return GoogleUserData{}, err
	}

	return userData, nil
}
