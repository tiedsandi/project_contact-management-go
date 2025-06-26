package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/models"
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

	contact := &models.Contact{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	savedContact, err := services.CreateContact(userId, contact)
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

func SearchContacts(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		return
	}
	userId := userIdInterface.(uint)

	name := c.Query("name")
	email := c.Query("email")
	phone := c.Query("phone")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	contacts, total, err := services.SearchContacts(userId, name, email, phone, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	var contactResponses []response.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, response.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email,
			Phone:     contact.Phone,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(size)))

	c.JSON(http.StatusOK, gin.H{
		"data": contactResponses,
		"paging": gin.H{
			"page":       page,
			"total_page": totalPages,
			"total_item": total,
		},
	})
}

func GetContact(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		return
	}
	userId := userIdInterface.(uint)

	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil || contactId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contact ID"})
		return
	}

	contact, err := services.GetContact(userId, uint(contactId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	resp := response.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func UpdateContact(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		return
	}
	userId := userIdInterface.(uint)

	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil || contactId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contact ID"})
		return
	}

	var req request.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	contact := &models.Contact{
		ID:        uint(contactId),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	savedContact, err := services.UpdateContact(userId, uint(contactId), contact)
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

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func DeleteContact(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		return
	}
	userId := userIdInterface.(uint)

	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil || contactId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contact ID"})
		return
	}

	err = services.DeleteContact(userId, uint(contactId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
