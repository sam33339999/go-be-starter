package routes

import (
	"github.com/sam33339999/go-be-starter/api/controllers"
	"github.com/sam33339999/go-be-starter/lib"
)

type AuthRoutes struct {
	logger     lib.Logger
	handler    lib.RequestHandler
	controller controllers.AuthController
}

func (s AuthRoutes) Setup() {
	s.logger.Debug("[Auth Controller] Setting up auth routes")

	api := s.handler.Gin.Group("/api")
	{
		api.GET("/login", s.controller.Login)
	}
}

func NewAuthRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
) AuthRoutes {
	return AuthRoutes{
		logger:  logger,
		handler: handler,
	}
}
