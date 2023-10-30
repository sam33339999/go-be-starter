package domains

import (
	"github.com/sam33339999/go-be-starter/models"
)

type AuthService interface {
	Authorize(token string) (bool, error)
	CreateToken(models.User) string
}
