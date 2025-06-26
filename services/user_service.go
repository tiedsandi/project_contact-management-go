package services

import (
	"errors"
	"strings"

	"github.com/tiedsandi/project_contact-management-go/models"
	"github.com/tiedsandi/project_contact-management-go/repositories"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func CreateUser(user *models.User) error {
	// Validation level service
	if user.Username == "" || strings.Contains(user.Username, " ") {
		return errors.New("username cannot be empty or contain spaces")
	}

	if !utils.IsValidPassword(user.Password) {
		return errors.New("password must be at least 6 characters long and contain both letters and numbers")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	user.Password, _ = utils.HashPassword(user.Password)

	err := repositories.CreateUser(user)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return errors.New("username already used")
	}

	return err
}

func Login(username, password string) (*models.User, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("username or password wrong")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("username or password wrong")
	}

	return user, nil
}

func UpdateUserByID(userID uint, name *string, password *string) (*models.User, error) {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if name != nil {
		if len(*name) > 100 {
			return nil, errors.New("name length max 100")
		}
		user.Name = *name
	}

	if password != nil {
		if !utils.IsValidPassword(*password) {
			return nil, errors.New("password must be at least 6 characters long and contain both letters and numbers")
		}
		hashed, _ := utils.HashPassword(*password)
		user.Password = hashed
	}

	err = repositories.UpdateUser(user)
	return user, err
}
