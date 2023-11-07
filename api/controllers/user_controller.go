package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sam33339999/go-be-starter/lib"
)

type UserController struct {
	logger lib.Logger
	// userService domains.UserService
}

func NewUserController(
	logger lib.Logger,
	// userService domains.UserService,
) UserController {
	return UserController{
		logger: logger,
		// userService: userService,
	}
}

func (u UserController) GetUsers(c *gin.Context) {
}
