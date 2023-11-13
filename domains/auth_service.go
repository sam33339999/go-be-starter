package domains

import "github.com/sam33339999/go-be-starter/models"

type AuthService interface {
	// middleware check token valid
	Authorize(token string) (bool, error)
	// for mock login
	Attempt(id uint) (string, error)
	// create jwt token
	CreateToken(models.User) (string, string)
	// refresh jwt token
	RefreshToken(refreshToken string) (string, string, error)
}
