package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tiedsandi/project_contact-management-go/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
}

func (u *User) Save(db *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = db.Create(u).Error

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errors.New("username already used")
		}
		return err
	}

	return nil
}
