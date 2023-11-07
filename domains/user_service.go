package domains

import "github.com/sam33339999/go-be-starter/models"

type UserService interface {
	// WithTrx(trxHandle *gorm.DB) UserService
	GetOneUser(id uint) (models.User, error)
	// GetAllUser() ([]models.User, error)

	// CreateUser(models.User) error
	// UpdateUser(models.User) error
	// DeleteUser(models.User) error
}
