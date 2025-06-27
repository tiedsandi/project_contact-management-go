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

func UpdateAddressForContact(userID uint, contactID uint, addressID uint, req *models.Address) (*models.Address, error) {
	address, err := repositories.GetAddressByID(addressID)
	if err != nil {
		return nil, err
	}

	if address.UserID != userID || address.ContactId != contactID {
		return nil, repositories.ErrNotFound()
	}

	address.Street = req.Street
	address.City = req.City
	address.Province = req.Province
	address.Country = req.Country
	address.PostalCode = req.PostalCode

	err = repositories.UpdateAddress(address)
	return address, err
}
