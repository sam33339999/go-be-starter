package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sam33339999/go-be-starter/domains"
	"github.com/sam33339999/go-be-starter/lib"
)

type JWTAuthController struct {
	logger  lib.Logger
	service domains.AuthService
	// userService domains.UserService
}

func NewJwtAuthController(
	logger lib.Logger,
	service domains.AuthService,
	// userService domains.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:  logger,
		service: service,
		// userService: userService,
	}
}

func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Debug("[Auth Controller] Registering new user")

	// user := jwt.userService.GetOneUser(uint(1))
	// token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "register route",
		"user":    "Hello World",
	})
}

func (jwt JWTAuthController) Login(c *gin.Context) {
	jwt.logger.Debug("[Auth Controller] Logging in user")

	token := jwt.service.CreateToken()

	c.JSON(200, gin.H{
		"token": token,
	})
}

func (jwt JWTAuthController) MockLogin(c *gin.Context) {
	jwt.logger.Debug("[Auth Controller] Logging in user")

	token := jwt.service.Attempt(1)
	c.JSON(200, gin.H{
		"attempt": 1,
		"token":   token,
	})
}
