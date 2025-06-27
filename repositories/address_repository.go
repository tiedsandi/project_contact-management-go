package repositories

import (
	"github.com/tiedsandi/project_contact-management-go/config"
	"github.com/tiedsandi/project_contact-management-go/models"
)

func CreateAddress(address *models.Address) error {
	return config.DB.Create(address).Error
}
