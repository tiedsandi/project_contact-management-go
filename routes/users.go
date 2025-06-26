package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/controllers"
	"github.com/tiedsandi/project_contact-management-go/middlewares"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/api/users/signup", controllers.Signup)
	router.POST("/api/users/login", controllers.Login)

	userPrivate := router.Group("/api/users")
	userPrivate.Use(middlewares.Authenticate)
	userPrivate.GET("/profile", controllers.GetUser)
	userPrivate.PUT("/profile", controllers.UpdateUser)
}
