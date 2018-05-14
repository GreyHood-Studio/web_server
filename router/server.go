package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GameServer struct {
	// [ip]:[port]
	Address			string
	MaxRoom			int
	CurrentRoom		int
}

type NPCServer struct {
	Address			string
}

var GameServers map[string]GameServer
var NPCServers map[string]NPCServer

func setServerRoute(router *gin.Engine) {

	server := router.Group("/server")
	{
		server.POST("/game", registGameServer)
		server.POST("/npc", registNPCServer)
		server.DELETE("/game", removeGameServer)
		server.DELETE("/npc", removeNPCServer)
	}
}

func registGameServer(c *gin.Context)  {
	var form GameServer
	if err := c.ShouldBind(&form); err == nil {
		c.JSON(http.StatusOK, gin.H{"serverId": "test"})
	}
}

func registNPCServer(c *gin.Context)  {
	var form NPCServer
	if err := c.ShouldBind(&form); err == nil {

	}
}

func removeGameServer(c *gin.Context)  {
	c.PostForm("serverID")
}

func removeNPCServer(c *gin.Context)  {
	c.PostForm("serverID")
}