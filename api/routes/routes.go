package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	userRoutes UserRoutes,
	authRoutes AuthRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
