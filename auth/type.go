package auth

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type ClaimsType struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
