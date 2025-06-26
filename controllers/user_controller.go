package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/requests"
	"github.com/tiedsandi/project_contact-management-go/responses"
	"github.com/tiedsandi/project_contact-management-go/services"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func Signup(c *gin.Context) {
	var req requests.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
	}

	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
		return
	}

	user, err := services.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Username, user.Name)

	c.JSON(http.StatusOK, gin.H{"data": responses.TokenResponse{Token: token}})
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

	var req requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
		return
	}

	user, err := services.UpdateUserByID(userId, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}})
}
