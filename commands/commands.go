package commands

import (
	"context"

	"github.com/sam33339999/go-be-starter/lib"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var cmds = map[string]lib.Command{
	"app:serve": NewServeCommand(),
	"jwt:gen":   NewJwtTokenGenCommand(),
}

func GetSubCommands(opt fx.Option) []*cobra.Command {
	var subCommands []*cobra.Command

	for name, cmd := range cmds {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

func WrapSubCommand(
	name string,
	cmd lib.Command,
	opt fx.Option,
) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			logger := lib.GetLogger()
			opts := fx.Options(
				fx.WithLogger(func() fxevent.Logger {
					return logger.GetFxLogger()
				}),
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer app.Stop(ctx)

			if err != nil {
				logger.Panic(err)
			}
		},
	}
	cmd.SetUp(wrappedCmd)
	return wrappedCmd
}
