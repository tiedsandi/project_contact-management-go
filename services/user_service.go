package services

import (
	"errors"

	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func RegisterUser(user *models.User) error {
	if user.Username == "" || user.Name == "" {
		return errors.New("username and name are required")
	}

	if !utils.IsValidPassword(user.Password) {
		return errors.New("invalid password format")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	return repositories.CreateUser(config.DB, user)
}

func AuthenticateUser(username, password string) (*models.User, error) {
	user, err := repositories.GetUserByUsername(config.DB, username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

func UpdateUser(userId uint, name *string, password *string) (*models.User, error) {
	user, err := repositories.GetUserByID(config.DB, userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if name != nil {
		user.Name = *name
	}

	if password != nil {
		if !utils.IsValidPassword(*password) {
			return nil, errors.New("invalid password format")
		}
		hashed, _ := utils.HashPassword(*password)
		user.Password = hashed
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
