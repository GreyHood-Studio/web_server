package api

import "github.com/gin-gonic/gin"

func SetAPIRoute(router *gin.Engine) {
	setBasicRoute(router)
	setLoginRoute(router)
}

func setBasicRoute(router *gin.Engine) {
	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}