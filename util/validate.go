package util

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// ProblemArgError - returned by ValidateSingleProblemArg when the problem argument is not parseable
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

// ValidateSingleProblemArg - Validate that a command has been called with a valid problem argument
func ValidateSingleProblemArg(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Requires a problem number argument (e.g. 42)")
	}
	return ValidateProblemStr(args[0])
}
