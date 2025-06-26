package services

import (
	"errors"

	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
)

func CreateContact(userId uint, contact *models.Contact) (*models.Contact, error) {
	if contact.FirstName == "" {
		return nil, errors.New("first name is required")
	}

	contact.UserID = userId
	err := repositories.CreateContact(contact)
	return contact, err
}

func SearchContacts(userId uint, name, email, phone string, page, size int) ([]models.Contact, int64, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * size

	contacts, total, err := repositories.SearchContacts(userId, name, email, phone, offset, size)
	return contacts, total, err
}
