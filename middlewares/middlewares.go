package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPanicMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	panicMiddleware PanicMiddleware,
) Middlewares {
	return Middlewares{
		panicMiddleware,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
