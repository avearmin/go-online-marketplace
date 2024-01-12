package auth

import "errors"

var (
	ErrInvalidSignature = errors.New("invalid signature")
	ErrTokenExpired     = errors.New("token expired")
	ErrInvalidIssuer    = errors.New("invalid issuer")
)
