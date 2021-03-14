package consts

import _ "embed"

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

// CLIName - default name of the eulercli executable
const CLIName = "eulercli"

// DefaultTemplatesDir - default relative path of the templates directory
const DefaultTemplatesDir = "eulercli-templates"

// DirPerm - permissions used for created directories
const DirPerm = 0755

// FilePerm - permissions used for created files
const FilePerm = 0655

// MissingAnswerMD5 - string used whenever a hashed answer cannot be found
const MissingAnswerMD5 = "?"

// ProblemsText - Embedded document with many Project Euler problems
//go:embed assets/project_euler_problems.txt
var ProblemsText string

// SolutionsTextURL - URL of some problem solutions collected by LuckyToilet
const SolutionsTextURL = "https://raw.githubusercontent.com/luckytoilet/projecteuler-solutions/master/Solutions.md"

// TempDirPath - path of a directory used to store temporary files
const TempDirPath = "/tmp/eulercli"

// TemplRepoName - name of the template github repository
const TemplRepoName = "eulercli-templates"

// TemplRepoOwner - owner of the template github repository
const TemplRepoOwner = "koomen"

// TemplRepoBranch - name of the template github master branch
const TemplRepoBranch = "main"

// ZippedRepoURL - URL for downloading a zipped copy of the eulercli repo
const ZippedRepoURL = "https://github.com/koomen/eulercli/archive/main.zip"
