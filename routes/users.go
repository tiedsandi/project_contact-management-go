package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save(config.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}
