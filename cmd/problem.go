package cmd

import (
	"fmt"
	"strconv"

	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(problemCmd)
}

// problemCmd
var problemCmd = &cobra.Command{
	Use:   "problem [problem]",
	Short: "Return the text of the specified problem",
	Long:  fmt.Sprintf(`Return the text of the specified problem.`),
	Args:  util.ValidateProblemArg,
	Run: func(cmd *cobra.Command, args []string) {
		problem, _ := strconv.Atoi(args[0])
		fmt.Println(util.GetProblemText(problem))
	},
}
