package middlewares

import (
	"github.com/sam33339999/go-be-starter/lib"
)

type CorsMiddleware struct {
	handler lib.RequestHandler
	logger  lib.Logger
	env     lib.Env
}

func NewCorsMiddleware(
	handler lib.RequestHandler,
	logger lib.Logger,
	env lib.Env,
) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

func (m CorsMiddleware) Setup() {

	m.logger.Info("Setting up cors middleware")

}
