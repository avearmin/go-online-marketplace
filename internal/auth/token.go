package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

func (tt TokenType) String() string {
	return string(tt)
}

const (
	AcessToken   TokenType = "gorage-sale-access"
	RefreshToken TokenType = "gorage-sale-refresh"
)

func CreateJWT(id uuid.UUID, jwtSecret string, expiresIn time.Duration, tokenType TokenType) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    tokenType.String(),
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
	return CreateJWT(id, jwtSecret, expiresIn, AcessToken)
}

func CreateRefreshToken(id uuid.UUID, jwtSecret string) (string, error) {
	expiresIn := (60 * 24) * time.Hour
	return CreateJWT(id, jwtSecret, expiresIn, RefreshToken)
}
