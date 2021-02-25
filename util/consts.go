package util

import _ "embed"

// DefaultLanguageFilename - name of the file in which the default language is stored
const DefaultLanguageFilename = ".eulerlang"

// ProblemTextURL - url for the raw problem text file
const ProblemTextURL = "https://raw.githubusercontent.com/davidcorbin/euler-offline/master/project_euler_problems.txt"

// ProblemSolutionText - text of problems and solutions
//go:embed project_euler_problems.txt
var ProblemSolutionText string
