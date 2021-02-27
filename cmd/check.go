package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/koomen/eulercli/consts"
	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

// checkCmd
var checkCmd = &cobra.Command{
	Use:   "check [problem] [answer]",
	Short: "Check the answer for the a specified problem",
	Long:  fmt.Sprintf(`Check the answer for the a specified problem.`),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Requires a problem number and answer arguments (e.g. 10 142913828922)")
		}
		return util.ValidateProblemStr(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		problemNum, _ := strconv.Atoi(args[0])
		guess := args[1]

		correctness, err := util.CheckAnswer(problemNum, guess)
		cobra.CheckErr(err)

		switch correctness {
		case consts.Correct:
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Printf("The answer %s for problem %d is %s\n", guess, problemNum, green("correct"))
		case consts.Incorrect:
			red := color.New(color.FgRed).SprintFunc()
			fmt.Printf("The answer %s for problem %d is %s\n", guess, problemNum, red("incorrect"))
		case consts.Unknown:
			fmt.Printf("The answer to problem %d is unknown\n", problemNum)
		}
	},
}
