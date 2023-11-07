package repositories

import "github.com/sam33339999/go-be-starter/lib"

type UserRepository struct {
	lib.Database
	logger lib.Logger
}

func NewUserRepository(
	db lib.Database,
	logger lib.Logger,
) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}
