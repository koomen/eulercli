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
	Args:  util.ValidateSingleProblemArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		problemNum, _ := strconv.Atoi(args[0])
		problem, err := util.GetProblem(problemNum)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", problem.ProblemText)
		return nil
	},
}
