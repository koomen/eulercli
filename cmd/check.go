package cmd

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"

	"github.com/fatih/color"
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
		problem, _ := strconv.Atoi(args[0])
		guess := args[1]
		hashedGuess := fmt.Sprintf("%x", md5.Sum([]byte(guess)))
		hashedAnswer := util.GetHashedAnswer(problem)

		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		if hashedGuess == hashedAnswer {
			fmt.Printf("The answer %s for problem %d is %s\n", guess, problem, green("correct"))
		} else {
			fmt.Printf("The answer %s for problem %d is %s\n", guess, problem, red("incorrect"))
		}

	},
}
