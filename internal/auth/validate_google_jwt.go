package auth

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	jwt.RegisteredClaims
}

func ValidateGoogleJWT(tokenString, pem, clientId string) (GoogleClaims, error) {
	var claimsStruct GoogleClaims

	_, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			return GoogleClaims{}, ErrTokenExpired
		}
		return GoogleClaims{}, err
	}

	if claimsStruct.Issuer != "accounts.google.com" && claimsStruct.Issuer != "https://accounts.google.com" {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	if ok := validateClaimsAudience(claimsStruct.Audience, clientId); !ok {
		return GoogleClaims{}, errors.New("aud is invalid")
	}

	return claimsStruct, nil
}

func validateClaimsAudience(claimsAudience jwt.ClaimStrings, clientId string) bool {
	for _, aud := range claimsAudience {
		if aud == clientId {
			return true
		}
	}
	return false
}
