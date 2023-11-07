package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sam33339999/go-be-starter/domains"
	"github.com/sam33339999/go-be-starter/lib"
)

type JWTAuthController struct {
	logger      lib.Logger
	service     domains.AuthService
	userService domains.UserService
}

func NewJwtAuthController(
	logger lib.Logger,
	service domains.AuthService,
	userService domains.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Debug("[Auth Controller] Registering new user")

	u, _ := jwt.userService.GetOneUser(uint(2))

	c.JSON(200, gin.H{
		"message": "register route",
		"user":    u,
	})
}

func (jwt JWTAuthController) Login(c *gin.Context) {
	jwt.logger.Debug("[Auth Controller] Logging in user")

	token, rToken := jwt.service.CreateToken()

	c.JSON(200, gin.H{
		"token":         token,
		"refresh_token": rToken,
	})
}

func (jwt JWTAuthController) MockLogin(c *gin.Context) {
	jwt.logger.Debug("[Auth Controller] mock logging in user")

	id := c.Query("id")

	jwt.logger.Debug("[Auth Controller] id: ", id)

	parseId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		jwt.logger.Error("[Auth Controller] Error parsing id: ", err)

		c.JSON(400, gin.H{
			"message": "id must be a number",
			"err":     err,
		})
		return
	}

	token, err := jwt.service.Attempt(uint(parseId))
	c.JSON(200, gin.H{
		"attempt": parseId,
		"token":   token,
	})
}
