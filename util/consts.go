package util

import _ "embed"

// DefaultLanguageFilename - name of the file in which the default language is stored
const DefaultLanguageFilename = ".eulerlang"

// ProblemTextURL - url for the raw problem text file
const ProblemTextURL = "https://raw.githubusercontent.com/davidcorbin/euler-offline/master/project_euler_problems.txt"

// MissingAnswerMD5 - string used whenever a hashed answer cannot be found
const MissingAnswerMD5 = "?"

// ProblemSolutionText - text of problems and solutions
//go:embed project_euler_problems.txt
var ProblemSolutionText string

// Correctness - type representing three states of "correctness": correct, incorrect, and unknown
type Correctness int

const (
	// Correct - used when a submitted guess is correct
	Correct Correctness = 1

	// Incorrect - used when a submitted guess is incorrect
	Incorrect Correctness = 0

	// Unknown - used when we don't have an answer for a given problem
	Unknown Correctness = -1
)
