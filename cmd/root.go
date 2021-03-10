package cmd

import (
	"fmt"

	"github.com/koomen/eulercli/consts"
	"github.com/spf13/cobra"
)

// CliName holds name of the CLI executable
var (
	Overwrite bool
	Verbose   bool
)

// init - Initialize the root command
func init() {

}

// RootCmd is cobra command executed when the CLI is called without any arguments
var rootCmd = &cobra.Command{
	Use:     consts.CLIName,
	Version: "0.2.0",
	Short:   fmt.Sprintf("%s is a CLI for working on Project Euler problems", consts.CLIName),
	Long: fmt.Sprintf(`%s is a CLI for working on Project Euler problems

Use it to create templated solutions to new problems, execute solutions, and check answers.`, consts.CLIName),
}

// Execute calls the eponymous function on the root command
func Execute() {
	rootCmd.Execute()
}
