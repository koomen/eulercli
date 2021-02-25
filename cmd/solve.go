package cmd

import (
	"fmt"

	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	solveCmd.MarkFlagRequired("language")
	rootCmd.AddCommand(solveCmd)
}

// solveCmd
var solveCmd = &cobra.Command{
	Use:   "solve [problem]",
	Short: "Execute solution for the specified problem",
	Long:  fmt.Sprintf(`Execute solution for the specified problem.`),
	Args:  util.ValidateProblemArg,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Solve command not implemented")
	},
}
