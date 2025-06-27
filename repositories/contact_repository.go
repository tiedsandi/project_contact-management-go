package repositories

import (
	"errors"

	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
	"gorm.io/gorm"
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

func GetContact(userId uint, contactId uint) (*models.Contact, error) {
	var contact models.Contact
	err := config.DB.Where("user_id = ? AND id = ?", userId, contactId).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func UpdateContact(contact *models.Contact) error {
	return config.DB.Save(contact).Error
}

func DeleteContactSoft(userId uint, contactId uint) error {
	return config.DB.Where("user_id = ? AND id = ?", userId, contactId).Delete(&models.Contact{}).Error
}

func DeleteContactHard(userId uint, contactId uint) error {
	return config.DB.Unscoped().Where("user_id = ? AND id = ?", userId, contactId).Delete(&models.Contact{}).Error
}

func GetContactByIDAndUserID(contactID uint, userID uint) (*models.Contact, error) {
	var contact models.Contact
	err := config.DB.Where("id = ? AND user_id = ?", contactID, userID).First(&contact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found or unauthorized")
		}
		return nil, err
	}
	return &contact, nil
}
