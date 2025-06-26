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

func GetContact(userId, contactId uint) (*models.Contact, error) {
	contact, err := repositories.GetContact(userId, contactId)
	if err != nil {
		return nil, err
	}
	if contact == nil {
		return nil, errors.New("contact not found")
	}
	return contact, nil
}

func UpdateContact(userId, contactId uint, contact *models.Contact) (*models.Contact, error) {
	if contact.FirstName == "" {
		return nil, errors.New("first name is required")
	}

	existingContact, err := repositories.GetContact(userId, contactId)
	if err != nil {
		return nil, err
	}
	if existingContact == nil {
		return nil, errors.New("contact not found")
	}

	existingContact.FirstName = contact.FirstName
	existingContact.LastName = contact.LastName
	existingContact.Email = contact.Email
	existingContact.Phone = contact.Phone

	err = repositories.UpdateContact(existingContact)
	return existingContact, err
}

func DeleteContact(userId, contactId uint) error {
	contact, err := repositories.GetContact(userId, contactId)
	if err != nil {
		return err
	}
	if contact == nil {
		return errors.New("contact not found")
	}

	return repositories.DeleteContactHard(userId, contactId)
}
