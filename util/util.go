package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// CheckError provides a simple mechanism for handling error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// GetProblemText - Return the text of the specified problem
func GetProblemText(problem int) string {
	problemDelimeter := regexp.MustCompile(`Problem [0-9]+\s*\=+`)
	split := problemDelimeter.Split(ProblemSolutionText, -1)

	leadingSpaces := regexp.MustCompile(`(?m:^ +)`)

	return leadingSpaces.ReplaceAllString(strings.Trim(split[problem], "\n"), "")
}

// GetHashedAnswer - Return the hashed answer for a given problem
func GetHashedAnswer(problem int) string {
	re := regexp.MustCompile(`Answer:\s+(\S+)`)
	answers := re.FindAllStringSubmatch(ProblemSolutionText, -1)

	return answers[problem-1][1]
}

// ProblemArgError - returned by ValidateProblemArg when the problem argument is not parseable
type ProblemArgError struct {
	arg string
}

// Error - Convert a ProblemArgError into a string
func (e *ProblemArgError) Error() string {
	return fmt.Sprintf("Problem arg \"%s\" is not an integer", e.arg)
}

// ValidateProblemStr - Validate that the given string can be parsed as an integer
func ValidateProblemStr(problemArgStr string) error {
	if _, err := strconv.Atoi(problemArgStr); err != nil {
		return &ProblemArgError{problemArgStr}
	}
	return nil
}

// ValidateProblemArg - Validate that a command has been called with a valid problem argument
func ValidateProblemArg(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Requires a problem number argument (e.g. 42)")
	}
	return ValidateProblemStr(args[0])
}
