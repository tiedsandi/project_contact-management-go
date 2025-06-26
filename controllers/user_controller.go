package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func Signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	// Manual validation
	if user.Username == "" || strings.Contains(user.Username, " ") {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Username cannot be empty or contain spaces"})
		return
	}

	if !utils.IsValidPassword(user.Password) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 6 characters long and contain both letters and numbers"})
		return
	}

	if user.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Name is required"})
		return
	}

	if err := user.Save(config.DB); err != nil {
		if err.Error() == "username already used" {
			context.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(context *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
		return
	}

	user, err := models.GetUserByUsername(config.DB, request.Username)

	// Console log user
	// fmt.Printf("User: %+v\n", user)

	if err != nil {
		if err.Error() == "user not found" {
			context.JSON(http.StatusUnauthorized, gin.H{"errors": "Username or password wrong"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
		}
		return
	}

	if !user.CheckPassword(request.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{"errors": "Username or password wrong"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": gin.H{"token": token}})
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

	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": "User not found"})
		return
	}

	if request.Name != nil {
		if len(*request.Name) > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Name length max 100"})
			return
		}
		user.Name = *request.Name
	}

	if request.Password != nil {
		if !utils.IsValidPassword(*request.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Password must be at least 6 characters long and contain both letters and numbers"})
			return
		}

		hashedPassword, _ := utils.HashPassword(*request.Password)
		user.Password = hashedPassword
	}

	// Simpan ke DB
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"username": user.Username,
			"name":     user.Name,
		},
	})
}
