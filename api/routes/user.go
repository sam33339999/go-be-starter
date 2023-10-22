package routes

import "github.com/sam33339999/go-be-starter/lib"

type UserRoutes struct {
	logger  lib.Logger
	handler lib.RequestHandler
}

func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
) UserRoutes {
	return UserRoutes{
		logger:  logger,
		handler: handler,
	}
}
