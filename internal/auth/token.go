package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	AccessIssuer  string = "gorage-sale-access"
	RefreshIssuer string = "gorage-sale-refresh"
)

func CreateJWT(id uuid.UUID, jwtSecret string, expiresIn time.Duration, issuer string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject:   id.String(),
	})
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func CreateAccessToken(id uuid.UUID, jwtSecret string) (string, error) {
	expiresIn := 1 * time.Hour
	return CreateJWT(id, jwtSecret, expiresIn, AccessIssuer)
}

func CreateRefreshToken(id uuid.UUID, jwtSecret string) (string, error) {
	expiresIn := (60 * 24) * time.Hour
	return CreateJWT(id, jwtSecret, expiresIn, RefreshIssuer)
}

func ValidateJWT(jwtString, jwtSecret string) (uuid.UUID, error) {
	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	idString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != AccessIssuer {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
