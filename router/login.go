package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/GreyHood-Studio/web_server/controller"
	"github.com/gin-contrib/sessions"
	"github.com/GreyHood-Studio/server_util/random"
	"fmt"
	"github.com/GreyHood-Studio/web_server/model"
)


func setLoginRoute(router *gin.Engine) {
	router.POST("/login", handleRequestLogin)
}

func handleRequestLogin(c *gin.Context)  {
	var form model.Account
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err == nil {
		nickname, exist := controller.LoginUser(form.ID, form.Password)
		if exist{
			// create session key
			userSession := sessions.Default(c)

			v := userSession.Get(nickname)
			if v != nil {
				fmt.Println(nickname, " already exist session: ", v)
			}
			sessionKey := random.CreateRandomString(16)
			userSession.Set(nickname, sessionKey)
			userSession.Save()
			c.JSON(http.StatusOK, gin.H{"status": nickname, "session": sessionKey})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}