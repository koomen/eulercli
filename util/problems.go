package util

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"

	"github.com/koomen/eulercli/consts"
)

// EulerProblem - a structured type representing a project euler problem
type EulerProblem struct {
	ProblemNum       int
	PaddedProblemNum string
	ProblemText      string
	AnswerMD5        string
	Answer           string
}

// MissingProblemError - used when the user requests a problem that doesn't exist
type MissingProblemError struct {
	ProblemNum                 int
	LargestSupportedProblemNum int
}

// Error - convert a MissingProblemError into a string
func (e *MissingProblemError) Error() string {
	return fmt.Sprintf(
		"Problem %d not found. The euler CLI supports problems 1-%d.",
		e.ProblemNum, e.LargestSupportedProblemNum)
}

// getProblemText - extract problem text from consts/assets/project_euler_problems.txt
func getProblemText(problemNum int) (string, error) {
	re := regexp.MustCompile(`Problem [0-9]+\s*\=+`)
	split := re.Split(consts.ProblemsText, -1)
	largestSupportedProblemNum := len(split) - 1

	if problemNum > largestSupportedProblemNum || problemNum < 1 {
		return "", &MissingProblemError{problemNum, largestSupportedProblemNum}
	}

	rawText := split[problemNum]

	// Remove the MD5'd answer text
	re = regexp.MustCompile(`((?s:.*))Answer:\s+(\S+)\s*$`)
	submatch := re.FindStringSubmatch(rawText)
	if len(submatch) < 3 {
		return rawText, nil
	}

	// Remove leading spaces from problem text
	leadingSpaces := regexp.MustCompile(`(?m:^   )`)
	formattedProblemText := leadingSpaces.ReplaceAllString(strings.TrimSpace(submatch[1]), "")

	return formattedProblemText, nil
}

func getAnswer(problemNum int) (string, error) {
	re := regexp.MustCompile(`(?m:^[0-9]+.\s+)`)
	split := re.Split(consts.SolutionsText, -1)
	largestSupportedProblemNum := len(split) - 1

	if problemNum > largestSupportedProblemNum || problemNum < 1 {
		return "", &MissingProblemError{problemNum, largestSupportedProblemNum}
	}

	return strings.TrimSpace(split[problemNum]), nil

}

// GetProblem - return an EulerProblem instance corresponding to the given problem number
func GetProblem(problemNum int) (*EulerProblem, error) {

	text, probErr := getProblemText(problemNum)
	answer, ansErr := getAnswer(problemNum)

	if probErr != nil && ansErr != nil {
		return nil, ansErr
	}

	return &EulerProblem{
		ProblemNum:       problemNum,
		PaddedProblemNum: fmt.Sprintf("%04d", problemNum),
		ProblemText:      text,
		AnswerMD5:        fmt.Sprintf("%x", md5.Sum([]byte(answer))),
		Answer:           answer,
	}, nil
}
