package cmd

import (
	"fmt"
	"strconv"

	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(answerCmd)
}

// answerCmd
var answerCmd = &cobra.Command{
	Use:   "answer [problem]",
	Short: "Return the answer (hashed) to the specified problem",
	Long:  fmt.Sprintf(`Return the answer (hashed) to the specified problem.`),
	Args:  util.ValidateSingleProblemArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		problemNum, _ := strconv.Atoi(args[0])
		problem, err := util.GetProblem(problemNum)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", problem.AnswerMD5)
		return err
	},
}
