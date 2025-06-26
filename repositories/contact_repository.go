package repositories

import (
	"github.com/tiedsandi/project_contact-management-go/models"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) error
	Search(userId uint, name, email, phone string, offset, limit int) ([]models.Contact, int64, error)
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

func (r *contactRepository) Search(userId uint, name, email, phone string, offset, limit int) ([]models.Contact, int64, error) {
	var contacts []models.Contact
	var total int64

	query := r.db.Model(&models.Contact{}).Where("user_id = ?", userId)

	if name != "" {
		query = query.Where("first_name LIKE ? OR last_name LIKE ?", "%"+name+"%", "%"+name+"%")
	}

	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}

	// Hitung total data
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	if err := query.Offset(offset).Limit(limit).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}
