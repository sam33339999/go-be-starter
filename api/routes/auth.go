package routes

import (
	"github.com/sam33339999/go-be-starter/api/controllers"
	"github.com/sam33339999/go-be-starter/lib"
)

type AuthRoutes struct {
	logger     lib.Logger
	handler    lib.RequestHandler
	controller controllers.JWTAuthController
}

func (s AuthRoutes) Setup() {
	s.logger.Debug("[Auth Controller] Setting up auth routes")

	api := s.handler.Gin.Group("/api")
	{
		api.GET("/login", s.controller.Login)
		api.GET("/register", s.controller.Register)
		api.GET("/refresh-token", s.controller.Register)
	}
}

func NewAuthRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.JWTAuthController,
) AuthRoutes {
	return AuthRoutes{
		logger:     logger,
		handler:    handler,
		controller: controller,
	}
}
