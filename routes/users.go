package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	if user.Username == "" || strings.Contains(user.Username, " ") {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Username cannot be empty or contain spaces"})
		return
	}

	if len(user.Password) < 6 || !utils.HasLetter(user.Password) || !utils.HasNumber(user.Password) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 6 characters long and contain both letters and numbers"})
		return
	}

	if user.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Name is required"})
		return
	}

	err = user.Save(config.DB)

	if err != nil {
		if err.Error() == "username already used" {
			context.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}
