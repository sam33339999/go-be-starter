package domains

import (
	"github.com/sam33339999/go-be-starter/lib"
)

type UserService struct{}

func NewUserService(
	logger lib.Logger,
) UserService {
	return UserService{}
}
