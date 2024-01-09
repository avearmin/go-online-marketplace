package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type timeNowFunc func() time.Time

const (
	AccessIssuer  string = "gorage-sale-access"
	RefreshIssuer string = "gorage-sale-refresh"
)

func createJWT(id uuid.UUID, jwtSecret string, expiresIn time.Duration, issuer string, nowFunc timeNowFunc) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(nowFunc()),
		ExpiresAt: jwt.NewNumericDate(nowFunc().Add(expiresIn)),
		Subject:   id.String(),
	})
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func CreateAccessToken(id uuid.UUID, jwtSecret string) (string, error) {
	expiresIn := calcExpiry(1)
	return createJWT(id, jwtSecret, expiresIn, AccessIssuer, time.Now)
}

func CreateRefreshToken(id uuid.UUID, jwtSecret string) (string, error) {
	twoMonthsInHours := 1440
	expiresIn := calcExpiry(twoMonthsInHours)
	return createJWT(id, jwtSecret, expiresIn, RefreshIssuer, time.Now)
}

func calcExpiry(hours int) time.Duration {
	return time.Duration(hours) * time.Hour
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
