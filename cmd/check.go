package cmd

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

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
	Short: "Check the answer supplied on the command line or piped from a solution program",
	Long: ("Check the answer for the a specified problem.  If the answer is not\n" +
		"specified, eulercli will scan stdin for the correct answer. This means you\n" +
		"can pipe the output of your solution program to eulercli check, e.g.:\n\n" +
		"     julia solution.jl | eulercli check 42\n\n" +
		"if neither the answer or the problem number are specified, eulercli will \n" +
		"scan stdin for both arguments, matching e.g. \"problem 25\" or \"Problem 25\"\n" +
		"for problem 25"),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 2 {
			return errors.New("check takes at most two arguments")
		}
		if len(args) >= 1 {
			return util.ValidateProblemStr(args[0])
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var problemNum int
		var answer, input string

		// Store command line arguments (if present) in the appropriate variables
		if len(args) >= 1 {
			problemNum, _ = strconv.Atoi(args[0])
			if len(args) == 2 {
				answer = args[1]
			}
		}

		// Extract missing arguments from stdin
		if len(args) <= 1 {
			if len(args) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "Scanning stdin for problem number and correct answer...\n")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "Scanning stdin for correct answer...\n")
			}

			fmt.Fprintf(cmd.OutOrStdout(), "-------------------------------------------------------------------------------\n\n\n")
			buf := make([]byte, 1)
			var n int
			var err error = nil
			for err == nil {
				n, err = cmd.InOrStdin().Read(buf)
				if n > 0 {
					input += string(buf[0:n])
					cmd.OutOrStdout().Write(buf[0:n])
				}
			}
			if err != io.EOF {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "\n\n-------------------------------------------------------------------------------\n")

			// extract problemNum from input
			if len(args) == 0 {
				re := regexp.MustCompile(`[Pp]roblem\s*(\d+)`)
				submatch := re.FindAllStringSubmatch(input, 1)
				if submatch != nil && len(submatch) > 0 {
					problemNum, err = strconv.Atoi(submatch[0][1])
					if err != nil {
						return err
					}
					fmt.Fprintf(cmd.OutOrStdout(), "Extracted problem number %d from input\n", problemNum)
				} else {
					return errors.New("unable to extract problem number from input")
				}
			}
		}

		problem, err := util.GetProblem(problemNum)
		if err != nil {
			return err
		}

		if problem.Answer == "" {
			fmt.Fprintf(cmd.OutOrStdout(), "The correct answer for problem %d is unknown\n", problemNum)
			return nil
		}

		green := color.New(color.FgGreen).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		if answer != "" {
			if answer == problem.Answer {
				fmt.Fprintf(cmd.OutOrStdout(), green("Congratulations, %s is the correct answer to problem %d!\n"), answer, problemNum)
				return nil
			}
			fmt.Fprintf(cmd.OutOrStdout(), red("%s is not the correct answer to problem %d. Keep trying!\n"), answer, problemNum)
			return nil
		}

		if strings.Contains(input, problem.Answer) {
			fmt.Fprintf(cmd.OutOrStdout(), "Detected answer %s in input.\n", problem.Answer)
			fmt.Fprintf(cmd.OutOrStdout(), green("Congratulations, this is the correct answer to problem %d!\n"), problemNum)
			return nil
		}

		fmt.Fprintf(cmd.OutOrStdout(), red("Failed to find correct answer for problem %d in input. Keep trying!\n"), problemNum)
		return nil
	},
}
