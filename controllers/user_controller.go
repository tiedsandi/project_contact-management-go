package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/services"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid JSON format"})
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
		return
	}

	user, err := services.AuthenticateUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": token}})
}

func GetUser(c *gin.Context) {
	userId, _ := c.Get("userId")
	username, _ := c.Get("username")
	name, _ := c.Get("name")

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"userId":   userId,
			"username": username,
			"name":     name,
		},
	})
}

func UpdateUser(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		return
	}
	userId := userIdInterface.(uint)

	var request struct {
		Name     *string `json:"name"`
		Password *string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
		return
	}

	updatedUser, err := services.UpdateUser(userId, request.Name, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"username": updatedUser.Username,
			"name":     updatedUser.Name,
		},
	})
}
