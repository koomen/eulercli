package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/koomen/eulercli/consts"
	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

var (
	Language    string
	TemplateDir string
	DstDir      string
)

func init() {
	// Add a global language flag
	generateCmd.PersistentFlags().StringVarP(
		&Language,
		"language",
		"l",
		"",
		fmt.Sprintf(
			"specifies which templates in %s should be used (e.g. julia, golang)",
			consts.DefaultTemplatesDir,
		),
	)
	generateCmd.PersistentFlags().StringVarP(
		&TemplateDir,
		"template",
		"t",
		"",
		fmt.Sprintf(
			"indicates where %s should look for solution program templates (defaults to %s/<language>)",
			consts.CLIName,
			consts.DefaultTemplatesDir,
		),
	)
	generateCmd.PersistentFlags().StringVarP(
		&DstDir,
		"destination",
		"d",
		"",
		fmt.Sprintf(
			"indicates where %s should save generated solution files (defaults to ./<language>)",
			consts.CLIName,
		),
	)

	rootCmd.AddCommand(generateCmd)
}

// generateCmd
var generateCmd = &cobra.Command{
	Use:   "generate [problem]",
	Short: "Create a template solution for the specified problem",
	Long:  fmt.Sprintf(`Create a template solution for the specified problem`),
	Args:  util.ValidateSingleProblemArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		problemNum, _ := strconv.Atoi(args[0])
		problem, err := util.GetProblem(problemNum)
		if err != nil {
			return err
		}

		// Validate that we know where to look for template files
		if TemplateDir == "" {
			if Language == "" {
				return errors.New("language or template directory must be specified")
			}
			TemplateDir = filepath.Join(consts.DefaultTemplatesDir, strings.ToLower(Language))
		}

		// Validate that the templates directory exists
		_, err = os.Stat(TemplateDir)
		if os.IsNotExist(err) {
			if filepath.HasPrefix(TemplateDir, consts.DefaultTemplatesDir) {
				fmt.Fprintf(cmd.OutOrStdout(),
					"Template directory %s does not exist. You can use"+
						"\n\n     %s pull\n\n"+
						"to download or update default templates in %s.\n",
					TemplateDir,
					consts.CLIName,
					consts.DefaultTemplatesDir,
				)
				confirm, err := util.Confirm(
					"Would you like to do this now?",
					true,
					cmd.InOrStdin(),
					cmd.OutOrStdout(),
				)
				if err != nil {
					return err
				}
				if confirm {
					err = pullCmd.RunE(pullCmd, []string{})
					if err != nil {
						return err
					}
				}
			}

			// If the user decided to pull templates, this may have fixed the problem,
			// so we double-check before returning an error
			_, err = os.Stat(TemplateDir)
			if os.IsNotExist(err) {
				return fmt.Errorf("template directory %s does not exist", TemplateDir)
			}
		}

		// Validate that we know where to save rendered template files
		if DstDir == "" {
			if Language == "" {
				return errors.New("language or destination directory must be specified")
			}
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			DstDir = filepath.Join(wd, Language)
		}

		// Make sure that the destination directory isn't in the templates directory
		if filepath.HasPrefix(DstDir, TemplateDir) {
			return fmt.Errorf(
				"destination directory %s is inside the template directory %s",
				DstDir,
				TemplateDir,
			)
		}

		// Create a temporary destination directory in /tmp
		util.CreateTempDir()
		defer util.RemoveTempDir()
		tempDstDir := util.TempPath("rendered")
		err = os.MkdirAll(tempDstDir, consts.DirPerm)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Rendering templates from %s to %s\n", TemplateDir, DstDir)
		err = util.RenderTemplateDir(TemplateDir, tempDstDir, problem, true, cmd.InOrStdin(), cmd.OutOrStdout())
		if err != nil {
			return err
		}
		return util.SyncDirs(tempDstDir, DstDir, false, cmd.InOrStdin())
	},
}
