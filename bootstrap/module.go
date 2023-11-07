package bootstrap

import (
	"github.com/sam33339999/go-be-starter/api/controllers"
	"github.com/sam33339999/go-be-starter/api/routes"
	"github.com/sam33339999/go-be-starter/lib"
	"github.com/sam33339999/go-be-starter/middlewares"
	"github.com/sam33339999/go-be-starter/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	routes.Module,
	controllers.Module,
	middlewares.Module,
	services.Module,
)
