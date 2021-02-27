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
	ProblemNum  int
	ProblemText string
	AnswerMD5   string
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

// getProblemAndAnswerTextFromFile - extracts problem text from project_euler_problems.txt
func getProblemAndAnswerTextFromFile(problemNum int) (string, error) {
	re := regexp.MustCompile(`Problem [0-9]+\s*\=+`)
	split := re.Split(consts.ProblemSolutionText, -1)

	largestSupportedProblemNum := len(split) + 1

	if problemNum > largestSupportedProblemNum {
		return "", &MissingProblemError{problemNum, largestSupportedProblemNum}
	}

	return split[problemNum], nil
}

// extractProblemFromText - extracts the problem text from the combined problem/answer text
func extractProblemAndAnswerFromText(text string) (string, string) {
	re := regexp.MustCompile(`((?s:.*))Answer:\s+(\S+)\s*$`)
	submatch := re.FindStringSubmatch(text)

	if len(submatch) < 3 {
		return "", ""
	}

	// Extract and format problem text
	rawProblem := submatch[1]
	leadingSpaces := regexp.MustCompile(`(?m:^   )`)
	problem := leadingSpaces.ReplaceAllString(strings.TrimSpace(rawProblem), "")

	// Extract hashed answer
	answer := submatch[2]

	return problem, answer
}

// GetProblem - return an EulerProblem instance corresponding to the given problem number
func GetProblem(problemNum int) (*EulerProblem, error) {

	text, err := getProblemAndAnswerTextFromFile(problemNum)

	if err != nil {
		return nil, err
	}

	problem, answer := extractProblemAndAnswerFromText(text)

	return &EulerProblem{
		ProblemNum:  problemNum,
		ProblemText: problem,
		AnswerMD5:   answer,
	}, nil
}

// CheckAnswer - compares a guess against the answer to a question
func CheckAnswer(problemNum int, guess string) (consts.Correctness, error) {
	problem, err := GetProblem(problemNum)
	if err != nil {
		return consts.Unknown, err
	}

	if problem.AnswerMD5 == consts.MissingAnswerMD5 {
		return consts.Unknown, nil
	}

	hashedGuess := fmt.Sprintf("%x", md5.Sum([]byte(guess)))

	if hashedGuess == problem.AnswerMD5 {
		return consts.Correct, nil
	}

	return consts.Incorrect, nil
}
