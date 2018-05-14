package router

import "github.com/gin-gonic/gin"

func init()  {
	GameServers = make(map[string]GameServer)
	NPCServers = make(map[string]NPCServer)
}

func SetAPIRoute(router *gin.Engine) {
	setBasicRoute(router)
	setLoginRoute(router)
	setAccountRoute(router)
	setServerRoute(router)
	setRoomRoute(router)
}

func setBasicRoute(router *gin.Engine) {
	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}