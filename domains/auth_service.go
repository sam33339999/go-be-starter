package domains

import (
	"github.com/sam33339999/go-be-starter/lib"
)

type AuthService struct{}

func NewAuthService(
	logger lib.Logger,
) AuthService {
	return AuthService{}
}
