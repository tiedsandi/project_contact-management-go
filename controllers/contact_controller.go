package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
	"github.com/tiedsandi/project_contact-management-go/request"
	"github.com/tiedsandi/project_contact-management-go/response"
	"github.com/tiedsandi/project_contact-management-go/services"
)

func CreateContact(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		return
	}
	userId := userIdInterface.(uint)

	var req request.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	repo := repositories.NewContactRepository(config.DB)
	service := services.NewContactService(repo)

	contact := &models.Contact{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		UserID:    userId,
	}

	savedContact, err := service.CreateContact(userId, contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	resp := response.ContactResponse{
		ID:        savedContact.ID,
		FirstName: savedContact.FirstName,
		LastName:  savedContact.LastName,
		Email:     savedContact.Email,
		Phone:     savedContact.Phone,
	}

	c.JSON(http.StatusCreated, gin.H{"data": resp})
}
