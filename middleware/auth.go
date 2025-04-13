package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/syu6noob/go-nas/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": "Missing or invalid Authorization header",
				})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &auth.ClaimsType{}

		token, err := jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return auth.Variables().SecretKey, nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
