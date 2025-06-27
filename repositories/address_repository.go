package repositories

import (
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
)

func CreateAddress(address *models.Address) error {
	return config.DB.Create(address).Error
}

func GetAddressesByContactIDAndUserID(contactId, userId uint) ([]models.Address, error) {
	var addresses []models.Address
	err := config.DB.Where("contact_id = ? AND user_id = ?", contactId, userId).Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func GetAddressByIDAndUserID(addressId, userId uint) (*models.Address, error) {
	var address models.Address
	err := config.DB.Where("id = ? AND user_id = ?", addressId, userId).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func GetAddressByID(addressId uint) (*models.Address, error) {
	var address models.Address
	err := config.DB.First(&address, addressId).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func UpdateAddress(address *models.Address) error {
	return config.DB.Save(address).Error
}

func DeleteSoftAddressByIDAndUserID(addressId, userId uint) error {
	return config.DB.Where("id = ? AND user_id = ?", addressId, userId).Delete(&models.Address{}).Error
}

func DeleteHardAddressByIDAndUserID(addressId, userId uint) error {
	return config.DB.Unscoped().Where("id = ? AND user_id = ?", addressId, userId).Delete(&models.Address{}).Error
}

func ErrNotFound() error {
	return config.DB.Error
}

func DeleteAddressesByContactIdAndUserId(contactId, userId uint) error {
	return config.DB.Where("contact_id = ? AND user_id = ?", contactId, userId).Delete(&models.Address{}).Error
}

func DeleteAddressesByContactIdAndUserIdHard(contactId, userId uint) error {
	return config.DB.Unscoped().Where("contact_id = ? AND user_id = ?", contactId, userId).Delete(&models.Address{}).Error
}
