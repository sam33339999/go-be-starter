package bootstrap

import (
	"github.com/sam33339999/go-be-starter/api/controllers"
	"github.com/sam33339999/go-be-starter/api/routes"
	"github.com/sam33339999/go-be-starter/lib"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	routes.Module,
	controllers.Module,
)
