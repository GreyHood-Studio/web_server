package api

import "github.com/gin-gonic/gin"

var DB = make(map[string]string)

func setLoginRoute(router *gin.Engine) {

	router.GET("/user/:name", getUserID)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})
}

func getUserID(c *gin.Context) {
	// Get user value
	user := c.Params.ByName("name")
	value, ok := DB[user]
	if ok {
		c.JSON(200, gin.H{"user": user, "value": value})
	} else {
		c.JSON(200, gin.H{"user": user, "status": "no value"})
	}
}