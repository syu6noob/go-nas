package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generate(username string, duration time.Duration, key []byte) (string, error) {
	claims := &ClaimsType{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
