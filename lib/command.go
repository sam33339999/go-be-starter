package lib

import "github.com/spf13/cobra"

type CommandRunner interface{}

type Command interface {
	Short() string

	SetUp(cmd *cobra.Command)

	Run() CommandRunner
}
