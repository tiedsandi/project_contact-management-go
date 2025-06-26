package repositories

import (
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
)

func CreateContact(contact *models.Contact) error {
	return config.DB.Create(contact).Error
}

func SearchContacts(userId uint, name string, email string, phone string, offset int, size int) ([]models.Contact, int64, error) {
	var contacts []models.Contact
	var total int64

	query := config.DB.Model(&models.Contact{}).Where("user_id = ?", userId)

	if name != "" {
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+name+"%", "%"+name+"%")
	}
	if email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}
	if phone != "" {
		query = query.Where("phone ILIKE ?", "%"+phone+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(size).Find(&contacts).Error
	return contacts, total, err
}
