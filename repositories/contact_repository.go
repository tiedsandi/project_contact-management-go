package repositories

import (
	"github.com/tiedsandi/project_contact-management-go/models"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

func (r *contactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}
