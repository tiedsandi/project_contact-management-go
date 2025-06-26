package services

import (
	"errors"
	"strings"

	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
)

type ContactService interface {
	CreateContact(userId uint, contact *models.Contact) (*models.Contact, error)
	SearchContacts(userId uint, name, email, phone string, page, size int) ([]models.Contact, int64, error)
}

type contactService struct {
	repo repositories.ContactRepository
}

func NewContactService(repo repositories.ContactRepository) ContactService {
	return &contactService{repo}
}

func (s *contactService) CreateContact(userId uint, contact *models.Contact) (*models.Contact, error) {
	// Validasi sederhana
	if contact.FirstName == "" {
		return nil, errors.New("First name is required")
	}

	if !strings.Contains(contact.Email, "@") {
		return nil, errors.New("Email is not valid format")
	}

	// Inject userId ke contact
	contact.UserID = userId

	// Simpan
	if err := s.repo.Create(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *contactService) SearchContacts(userId uint, name, email, phone string, page, size int) ([]models.Contact, int64, error) {
	offset := (page - 1) * size
	return s.repo.Search(userId, name, email, phone, offset, size)
}
