package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/syu6noob/go-nas/api"
	"github.com/syu6noob/go-nas/middleware"
)

func notFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"Error": "Not found",
	})
}

func frontend(c *gin.Context) {
	c.File("./frontend/dist/index.html")
}

func main() {
	router := gin.Default()
	router.Use(
		middleware.ErrorMiddleware(),
		middleware.CorsMiddleware(),
	)

	api.Routes(router)

	router.Static("/assets", "./frontend/dist/assets")
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") || strings.HasPrefix(c.Request.URL.Path, "/static/") {
			notFound(c)
			return
		}
		frontend(c)
	})

	router.Run(":80")
}
