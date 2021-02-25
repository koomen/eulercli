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
	Args:  util.ValidateProblemArg,
	Run: func(cmd *cobra.Command, args []string) {
		problem, _ := strconv.Atoi(args[0])
		fmt.Println(util.GetHashedAnswer(problem))
	},
}
