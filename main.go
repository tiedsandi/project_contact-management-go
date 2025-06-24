package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/config"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}
