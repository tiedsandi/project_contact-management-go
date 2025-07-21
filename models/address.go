package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"id"`
	Street     string `gorm:"size:100;not null;uniqueIndex:idx_contact_street" json:"street"`
	City       string `gorm:"size:100" json:"city"`
	Province   string `gorm:"size:100;" json:"province"`
	Country    string `gorm:"size:20" json:"country"`
	PostalCode string `gorm:"not null" json:"postal_code"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ContactId uint    `gorm:"not null;uniqueIndex:idx_contact_street" json:"contact_id"`
	Contact   Contact `gorm:"foreignKey:ContactId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
