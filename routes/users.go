package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/controllers"
)

func UserRoutes(router *gin.Engine) {
	// Public routes (tanpa token)
	userPublic := router.Group("/api/users")
	{
		userPublic.POST("", controllers.Signup) // Register
		// userPublic.POST("/login", controllers.Login)        // Login
	}

	// Protected routes (butuh token / sudah login)
	// userPrivate := router.Group("/api/users")
	// userPrivate.Use(middlewares.Authenticate)
	// {
	// 	userPrivate.PATCH("/current", controllers.UpdateUser) // Update user info
	// 	userPrivate.GET("/current", controllers.GetUser)      // Get current user
	// 	userPrivate.DELETE("/logout", controllers.Logout)     // Logout
	// }
}
