package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.Migrate()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run()
}
