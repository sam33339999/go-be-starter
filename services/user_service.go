package services

import (
	"github.com/sam33339999/go-be-starter/domains"
	"github.com/sam33339999/go-be-starter/lib"
	"github.com/sam33339999/go-be-starter/models"
)

type UserService struct {
	env    lib.Env
	logger lib.Logger
}

func NewUserService(
	env lib.Env,
	logger lib.Logger,
) domains.UserService {
	return UserService{
		env:    env,
		logger: logger,
	}
}

func (s UserService) GetOneUser(id uint) (models.User, error) {
	return models.User{}, nil
}
