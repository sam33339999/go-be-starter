package routes

import (
	"github.com/sam33339999/go-be-starter/api/controllers"
	"github.com/sam33339999/go-be-starter/lib"
)

type UserRoutes struct {
	logger     lib.Logger
	handler    lib.RequestHandler
	controller controllers.UserController
}

func (s UserRoutes) Setup() {
	s.logger.Debug("[User Controller] Setting up user routes")

	api := s.handler.Gin.Group("/api")
	{
		api.GET("/users", s.controller.GetUsers)
	}
}

func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.UserController,
) UserRoutes {
	return UserRoutes{
		logger:     logger,
		handler:    handler,
		controller: controller,
	}
}
