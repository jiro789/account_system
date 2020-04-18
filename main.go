package main

import (
	"account_system/api/components/account"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	os.Setenv("PORT", "3000")
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./client/build", false)))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	account.Routes(api)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
