package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/GreyHood-Studio/web_server/controller"
	"github.com/GreyHood-Studio/web_server/model"
)

func setAccountRoute(router *gin.Engine) {

	account := router.Group("/account")
	{
		account.POST("/signup", handleRequestSignupForm)
		account.POST("/signupJson", handleRequestSignupJson)
		account.PUT("/password", handleRequestChangePasswordForm)
		account.PUT("/passwordJson", handleRequestChangePasswordJson)
		account.DELETE("/leave", handleRequestTerminateAccount)
	}
}

func handleRequestSignupForm(c *gin.Context)  {
	var form model.User
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err == nil {
		controller.SignUpAccount(form.ID, form.Password, form.NickName)
		c.JSON(http.StatusOK, gin.H{"status": "signup complete"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func handleRequestSignupJson(c *gin.Context)  {
	var json model.User
	if err := c.ShouldBindJSON(&json); err == nil {
		controller.SignUpAccount(json.ID, json.Password, json.NickName)
		c.JSON(http.StatusOK, gin.H{"status": "signup complete"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func handleRequestChangePasswordForm(c *gin.Context) {
	var form model.Account
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err == nil {
		controller.ChangePassword(form.ID, form.Password)
		c.JSON(http.StatusOK, gin.H{"status": "change password complete"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func handleRequestChangePasswordJson(c *gin.Context) {
	var json model.Account
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&json); err == nil {
		controller.ChangePassword(json.ID, json.Password)
		c.JSON(http.StatusOK, gin.H{"status": "change password complete"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func handleRequestTerminateAccount(c *gin.Context)  {

}