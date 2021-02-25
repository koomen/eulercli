package cmd

import (
	"fmt"

	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	startCmd.MarkFlagRequired("language")
	rootCmd.AddCommand(startCmd)
}

// startCmd
var startCmd = &cobra.Command{
	Use:   "start [problem]",
	Short: "Create a template solution for the specified problem",
	Long:  fmt.Sprintf(`Create a template solution for the specified problem in the specified language.`),
	Args:  util.ValidateProblemArg,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Start command called with language %s\n", Language)
	},
}
