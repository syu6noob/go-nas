package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func Validate(tokenStr string, key []byte) (*ClaimsType, error) {
	claims := &ClaimsType{}
	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
