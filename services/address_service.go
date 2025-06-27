package services

import (
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
)

func CreateAddressForContact(userID uint, contactID uint, req *models.Address) (*models.Address, error) {
	req.UserID = userID
	req.ContactId = contactID
	err := repositories.CreateAddress(req)
	return req, err
}
