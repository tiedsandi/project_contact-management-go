package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/controllers"
	"github.com/tiedsandi/project_contact-management-go/middlewares"
)

func UserRoutes(router *gin.Engine) {
	userPublic := router.Group("/api/users")
	{
		userPublic.POST("/signup", controllers.Signup)
		userPublic.POST("/login", controllers.Login)
	}

	userPrivate := router.Group("/api/users")
	userPrivate.Use(middlewares.Authenticate)
	{
		userPrivate.PATCH("/current", controllers.UpdateUser)
		userPrivate.GET("/current", controllers.GetUser)
		// userPrivate.DELETE("/logout", controllers.Logout)
	}
}
