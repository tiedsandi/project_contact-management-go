package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/controllers"
	"github.com/tiedsandi/project_contact-management-go/middlewares"
)

func AddressRoutes(router *gin.Engine) {
	addressPrivate := router.Group("/api/contact/:contactId/addresses")
	addressPrivate.Use(middlewares.Authenticate)

	addressPrivate.POST("", controllers.CreateAddress)
	addressPrivate.PUT("/:id", controllers.UpdateAddress)
	addressPrivate.GET("/:id", controllers.GetAddress)
	addressPrivate.GET("", controllers.ListAddresses)
	addressPrivate.DELETE("/:id", controllers.DeleteAddress)
}
