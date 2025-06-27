package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
	"github.com/tiedsandi/project_contact-management-go/request"
	"github.com/tiedsandi/project_contact-management-go/response"
	"github.com/tiedsandi/project_contact-management-go/services"
)

func CreateAddress(c *gin.Context) {
	userId, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	contactIdParam := c.Param("contactId")
	contactIdUint, err := strconv.ParseUint(contactIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contactId"})
		return
	}

	// Validasi contactnya ada dan milik user
	_, err = ValidateUserContact(uint(contactIdUint), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	var req request.CreateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	address := &models.Address{
		Street:     req.Street,
		City:       req.City,
		Province:   req.Province,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	}

	savedAddress, err := services.CreateAddressForContact(userId, uint(contactIdUint), address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	resp := response.AddressResponse{
		ID:         savedAddress.ID,
		Street:     savedAddress.Street,
		City:       savedAddress.City,
		Province:   savedAddress.Province,
		Country:    savedAddress.Country,
		PostalCode: savedAddress.PostalCode,
	}

	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

func ListAddresses(c *gin.Context) {
	userId, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	contactIdParam := c.Param("contactId")
	contactIdUint, err := strconv.ParseUint(contactIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contactId"})
		return
	}

	// Validasi contactnya ada dan milik user
	_, err = ValidateUserContact(uint(contactIdUint), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	addresses, err := repositories.GetAddressesByContactIDAndUserID(uint(contactIdUint), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to retrieve addresses"})
		return
	}
	var addressesResponse []response.AddressResponse
	for _, address := range addresses {
		addressesResponse = append(addressesResponse, response.AddressResponse{
			ID:         address.ID,
			Street:     address.Street,
			City:       address.City,
			Province:   address.Province,
			Country:    address.Country,
			PostalCode: address.PostalCode,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": addressesResponse})
}

func GetAddress(c *gin.Context) {
	userId, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	contactIdParam := c.Param("contactId")
	contactIdUint, err := strconv.ParseUint(contactIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contactId"})
		return
	}

	addressIdParam := c.Param("addressId")
	addressIdUint, err := strconv.ParseUint(addressIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid addressId"})
		return
	}

	// Validasi contactnya ada dan milik user
	_, err = ValidateUserContact(uint(contactIdUint), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	address, err := repositories.GetAddressByIDAndUserID(uint(addressIdUint), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": "Address not found"})
		return
	}

	resp := response.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		Country:    address.Country,
		PostalCode: address.PostalCode,
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func UpdateAddress(c *gin.Context) {
	userId, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	contactIdParam := c.Param("contactId")
	contactIdUint, err := strconv.ParseUint(contactIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contactId"})
		return
	}

	addressIdParam := c.Param("addressId")
	addressIdUint, err := strconv.ParseUint(addressIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid addressId"})
		return
	}

	// Validasi contactnya ada dan milik user
	_, err = ValidateUserContact(uint(contactIdUint), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	var req request.UpdateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	address := &models.Address{
		ID:         uint(addressIdUint),
		Street:     req.Street,
		City:       req.City,
		Province:   req.Province,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	}

	address, err = services.UpdateAddressForContact(userId, uint(contactIdUint), uint(addressIdUint), address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": "Address not found or unauthorized"})
		return
	}

	resp := response.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		Country:    address.Country,
		PostalCode: address.PostalCode,
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func DeleteAddress(c *gin.Context) {
	userId, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	contactIdParam := c.Param("contactId")
	contactIdUint, err := strconv.ParseUint(contactIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid contactId"})
		return
	}

	addressIdParam := c.Param("addressId")
	addressIdUint, err := strconv.ParseUint(addressIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid addressId"})
		return
	}

	// Validasi contactnya ada dan milik user
	_, err = ValidateUserContact(uint(contactIdUint), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err = repositories.DeleteHardAddressByIDAndUserID(uint(addressIdUint), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": "Address not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		return 0, errors.New("unauthorized")
	}
	userId, ok := userIdInterface.(uint)
	if !ok {
		return 0, errors.New("invalid user id type")
	}
	return userId, nil
}

func ValidateUserContact(contactID uint, userID uint) (*models.Contact, error) {
	contact, err := repositories.GetContactByIDAndUserID(contactID, userID)
	if err != nil {
		return nil, errors.New("contact not found or unauthorized")
	}
	return contact, nil
}
