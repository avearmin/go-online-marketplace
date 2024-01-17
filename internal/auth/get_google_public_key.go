package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	if err = json.Unmarshal(dat, &myResp); err != nil {
		return "", err
	}

	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}

	return key, nil
}
