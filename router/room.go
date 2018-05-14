package router

import "github.com/gin-gonic/gin"

func setRoomRoute(router *gin.Engine)  {
	router.POST("/room", createRoom)
	router.POST("/room", changeRoom)
	router.DELETE("/room", removeRoom)
}

func createRoom(c *gin.Context)  {
	c.PostForm("serverID")
}

func changeRoom(c *gin.Context)  {
	c.PostForm("serverID")
}

func removeRoom(c *gin.Context)  {
	c.PostForm("serverID")
}