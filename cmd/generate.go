package cmd

import (
	"fmt"

	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

var Template string

func init() {
	generateCmd.MarkFlagRequired("language")
	generateCmd.LocalFlags().StringVarP(
		&Template,
		"template",
		"t",
		"",
		"programming language used to start and solve problems (default: julia)",
	)

	rootCmd.AddCommand(generateCmd)
}

// generateCmd
var generateCmd = &cobra.Command{
	Use:   "generate [problem]",
	Short: "Create a template solution for the specified problem",
	Long:  fmt.Sprintf(`Create a template solution for the specified problem in the specified language.`),
	Args:  util.ValidateSingleProblemArg,
	Run: func(cmd *cobra.Command, args []string) {
		// problemNum, _ := strconv.Atoi(args[0])
		// problem, err := util.GetProblem(problemNum)
		// cobra.CheckErr(err)

		// lang := strings.ToLower(Language)

	},
}
