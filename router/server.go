package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/GreyHood-Studio/web_server/model"
)


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
	var form model.GameServer
	if err := c.ShouldBind(&form); err == nil {
		c.JSON(http.StatusOK, gin.H{"serverId": "test"})
	}
}

func registNPCServer(c *gin.Context)  {
	var form model.NPCServer
	if err := c.ShouldBind(&form); err == nil {

	}
}

func removeGameServer(c *gin.Context)  {
	c.PostForm("serverID")
}

func removeNPCServer(c *gin.Context)  {
	c.PostForm("serverID")
}