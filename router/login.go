package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/GreyHood-Studio/web_server/controller"
	"github.com/GreyHood-Studio/server_util/random"
)


func setLoginRoute(router *gin.Engine) {
	router.POST("/login", handleRequestLogin)
}

func handleRequestLogin(c *gin.Context)  {
	var form Account
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err == nil {
		nickname, exist := controller.LoginUser(form.ID, form.Password)
		if exist{
			// create session key
			randkey := random.CreateRandomString(16)
			c.JSON(http.StatusOK, gin.H{"status": nickname, "session": randkey})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func passSessionKey()  {
	
}