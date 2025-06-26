package repositories

import (
	"errors"

	"github.com/tiedsandi/project_contact-management-go/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByID(db *gorm.DB, userId uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
