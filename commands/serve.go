package commands

import (
	"github.com/sam33339999/go-be-starter/api/routes"
	"github.com/sam33339999/go-be-starter/lib"
	"github.com/spf13/cobra"
)

type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) SetUp(cmd *cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		// middleware middlewares.Middlewares,
		env lib.Env,
		router lib.RequestHandler,
		logger lib.Logger,
		route routes.Routes,
		// database lib.Database,
	) {
		// middleware.Setup()
		route.Setup()

		logger.Info("Running server")
		_ = router.Gin.Run(":8888")
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
