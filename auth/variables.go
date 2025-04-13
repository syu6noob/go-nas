package auth

import (
	"os"
	"time"
)

type variablesType struct {
	Secret          string
	Refresh         string
	SecretKey       []byte
	RefreshKey      []byte
	Username        string
	Password        string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func Variables() variablesType {
	secret := os.Getenv("AUTH_SECRET")
	refresh := os.Getenv("AUTH_REFRESH")
	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	accessTokenTTL := 15 * time.Minute
	refreshTokenTTL := 24 * time.Hour

	return variablesType{
		secret,
		refresh,
		[]byte(secret),
		[]byte(refresh),
		username,
		password,
		accessTokenTTL,
		refreshTokenTTL,
	}
}
