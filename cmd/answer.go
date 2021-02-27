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
	Run: func(cmd *cobra.Command, args []string) {
		problemNum, _ := strconv.Atoi(args[0])
		problem, err := util.GetProblem(problemNum)
		cobra.CheckErr(err)
		fmt.Println(problem.AnswerMD5)
	},
}
