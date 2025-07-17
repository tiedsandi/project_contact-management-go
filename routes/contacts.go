package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/controllers"
	"github.com/tiedsandi/project_contact-management-go/middlewares"
)

func ContactRoutes(router *gin.Engine) {
	contactPrivate := router.Group("/api/contacts")
	contactPrivate.Use(middlewares.Authenticate)

	contactPrivate.POST("", controllers.CreateContact)
	contactPrivate.PUT("/:id", controllers.UpdateContact)
	contactPrivate.GET("/:id", controllers.GetContact)
	contactPrivate.GET("", controllers.SearchContacts)
	contactPrivate.DELETE("/:id", controllers.DeleteContact)
	contactPrivate.GET("/check-email", controllers.CheckEmailAvailable)

}
