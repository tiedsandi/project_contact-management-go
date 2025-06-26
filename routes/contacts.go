package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/controllers"
	"github.com/tiedsandi/project_contact-management-go/middlewares"
)

func ContactRoutes(router *gin.Engine) {
	contactPrivate := router.Group("/api/contacts")
	contactPrivate.Use(middlewares.Authenticate) // Biar wajib login / token

	contactPrivate.POST("", controllers.CreateContact)
	contactPrivate.PUT("/:id", controllers.UpdateContact)
	contactPrivate.GET("/:id", controllers.GetContact)
	contactPrivate.GET("", controllers.SearchContacts)
	contactPrivate.DELETE("/:id", controllers.DeleteContact)
}
