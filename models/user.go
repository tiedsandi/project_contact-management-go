package models

import (
	"github.com/tiedsandi/project_contact-management-go/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

func (u *User) Save(db *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return db.Create(u).Error
}
