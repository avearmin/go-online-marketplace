package auth

import (
	"errors"
	"strings"
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
	expiresIn, err := calcExpiry(1)
	if err != nil {
		return "", err
	}
	return createJWT(id, jwtSecret, expiresIn, AccessIssuer, time.Now)
}

func CreateRefreshToken(id uuid.UUID, jwtSecret string) (string, error) {
	twoMonthsInHours := 1440
	expiresIn, err := calcExpiry(twoMonthsInHours)
	if err != nil {
		return "", err
	}
	return createJWT(id, jwtSecret, expiresIn, RefreshIssuer, time.Now)
}

func calcExpiry(hours int) (time.Duration, error) {
	if hours < 0 {
		return 0, errors.New("negative hours")
	}
	return time.Duration(hours) * time.Hour, nil
}

func validateJWT(jwtString, jwtSecret, jwtIssuer string) (uuid.UUID, error) {
	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenSignatureInvalid.Error()) {
			return uuid.Nil, ErrInvalidSignature
		}
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			return uuid.Nil, ErrTokenExpired
		}
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
	if issuer != jwtIssuer {
		return uuid.Nil, ErrInvalidIssuer
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func ValidateAccessToken(jwtString, jwtSecret string) (uuid.UUID, error) {
	return validateJWT(jwtString, jwtSecret, AccessIssuer)
}

func ValidateRefreshToken(jwtString, jwtSecret string) (uuid.UUID, error) {
	return validateJWT(jwtString, jwtSecret, RefreshIssuer)
}
