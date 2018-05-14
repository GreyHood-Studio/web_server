package router

import (
	"github.com/gin-gonic/gin"
	"github.com/GreyHood-Studio/web_server/model"
)

func init()  {
	model.GameServers = make(map[string]model.GameServer)
	model.NPCServers = make(map[string]model.NPCServer)
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