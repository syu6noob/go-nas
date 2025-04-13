package api

import (
	"github.com/gin-gonic/gin"
	"github.com/syu6noob/go-nas/middleware"
)

func Routes(router *gin.Engine) {
	router.Use(
		middleware.CorsMiddleware(), // dev
	)

	router.POST("/login", Login)
	router.POST("/refresh", Refresh)

	api := router.Group("/api")
	{
		api.GET("/", Ping)
		api.GET("/open/*path", Stream)
		api.GET("/raw/*path", Download)
	}

	apiProtected := router.Group("/api")
	{
		apiProtected.Use(middleware.AuthMiddleware())
		apiProtected.GET("/info/*path", Info)
	}
}
