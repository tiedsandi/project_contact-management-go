package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100" json:"last_name"`
	Email     string `gorm:"size:100;uniqueIndex" json:"email"`
	Phone     string `gorm:"size:20" json:"phone"`
	UserID    uint   `gorm:"not null" json:"user_id"`
}
