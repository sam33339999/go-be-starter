package services

import (
	"github.com/sam33339999/go-be-starter/domains"
	"github.com/sam33339999/go-be-starter/lib"
	"github.com/sam33339999/go-be-starter/models"
)

type UserService struct {
	env    lib.Env
	logger lib.Logger
	db     lib.Database
}

func NewUserService(
	env lib.Env,
	logger lib.Logger,
	db lib.Database,
) domains.UserService {
	return UserService{
		env:    env,
		logger: logger,
		db:     db,
	}
}

func (s UserService) GetOneUser(id uint) (models.User, error) {
	var result models.User
	s.db.Model(models.User{ID: 10}).First(&result)

	return result, nil
}
