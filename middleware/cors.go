package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		godotenv.Load("./../.env")
		// host := os.Getenv("DEV_APP_HOST")

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Authorization, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
