package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/syu6noob/go-nas/auth"
)

type loginType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var creds loginType
	vars := auth.Variables()
	err := c.ShouldBindJSON(&creds)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// fmt.Printf("%s %s \n", creds.Username, creds.Password)
	// fmt.Printf("%s %s \n", vars.Username, vars.Password)

	if creds.Username != vars.Username || creds.Password != vars.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	accessToken, accessTokenErr := auth.Generate(
		creds.Username,
		vars.AccessTokenTTL,
		vars.SecretKey,
	)
	refreshToken, refreshTokenErr := auth.Generate(
		creds.Username,
		vars.RefreshTokenTTL,
		vars.RefreshKey,
	)
	if accessTokenErr != nil || refreshTokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token generation failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func Refresh(c *gin.Context) {
	vars := auth.Variables()

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	claims, claimsErr := auth.Validate(req.RefreshToken, auth.Variables().RefreshKey)
	if claimsErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	newAccessToken, newAccessTokenErr := auth.Generate(
		claims.Username,
		vars.AccessTokenTTL,
		vars.SecretKey,
	)

	if newAccessTokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
