package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
	"github.com/sam33339999/go-be-starter/lib"
)

type JwtCheckMiddleware struct {
	handler lib.RequestHandler
	logger  lib.Logger
	env     lib.Env
}

func NewJwtCheckMiddleware(
	handler lib.RequestHandler,
	logger lib.Logger,
	env lib.Env,
) JwtCheckMiddleware {
	return JwtCheckMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

func (m JwtCheckMiddleware) Setup() {
	m.logger.Info("Setting up jwt middleware")
	m.handler.Gin.Use(m.Handler())
}

func (m JwtCheckMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recover := recover(); recover != nil {
				switch value := recover.(type) {
				case eris.ErrLink:
					_ = value
				default:
					fmt.Printf("%+v", recover)

					errorMessage := fmt.Sprintf("%+v", recover)
					err := eris.New(errorMessage)

					eris := fmt.Sprintf("%v", eris.ToJSON(err, true))

					c.AbortWithStatusJSON(400, gin.H{
						"statusCode": 400,
						"details":    eris,
						"message":    errorMessage,
					})
				}
				return
			}
		}()
		c.Next()
	}
}
