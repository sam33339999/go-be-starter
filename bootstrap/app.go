package bootstrap

import (
	"github.com/sam33339999/go-be-starter/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   `go be starter`,
	Short: `Go BackendEnd Starter using Gin framework`,
	Long: `
	________                _____________________           _________ __                 __                
	/  _____/  ____          \______   \_   _____/          /   _____//  |______ ________/  |_  ___________ 
   /   \  ___ /  _ \   ______ |    |  _/|    __)_   ______  \_____  \\   __\__  \\_  __ \   __\/ __ \_  __ \
   \    \_\  (  <_> ) /_____/ |    |   \|        \ /_____/  /        \|  |  / __ \|  | \/|  | \  ___/|  | \/
	\______  /\____/          |______  /_______  /         /_______  /|__| (____  /__|   |__|  \___  >__|   
		   \/                        \/        \/                  \/           \/                 \/       
													 
	This is a command runner or cli for api architecture in golang. 
	Using this we can use underlying dependency injection container for running scripts. 
	Main advantage is that, we can use same services, repositories, infrastructure present in the application itself
	memo: http://patorjk.com/software/taag/ is used for generating ascii art 
	`,
	TraverseChildren: true,
}

type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
