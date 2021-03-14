package cmd

import (
	"fmt"

	"github.com/koomen/eulercli/consts"
	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	pullCmd.PersistentFlags().BoolVarP(
		&Overwrite,
		"overwrite",
		"o",
		false,
		"overwrite existing template or target files (default: false)",
	)
	pullCmd.PersistentFlags().Lookup("overwrite").NoOptDefVal = "true"

	pullCmd.PersistentFlags().BoolVarP(
		&Verbose,
		"verbose",
		"v",
		false,
		"verbose mode (default: false)",
	)
	pullCmd.PersistentFlags().Lookup("verbose").NoOptDefVal = "true"
	rootCmd.AddCommand(pullCmd)
}

// pullCmd
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: fmt.Sprintf("Download solution templates to %s", consts.DefaultTemplatesDir),
	Long:  fmt.Sprintf("Download solution templates to %s", consts.DefaultTemplatesDir),
	RunE: func(cmd *cobra.Command, args []string) error {
		util.CreateTempDir()
		defer util.RemoveTempDir()

		fmt.Fprintf(
			cmd.OutOrStdout(),
			"Downloading templates from %s\n",
			util.BuildRepoUrl(consts.TemplRepoOwner, consts.TemplRepoName),
		)

		zippedRepo := util.TempPath(fmt.Sprintf("%s-%s.zip", consts.TemplRepoName, consts.TemplRepoBranch))
		err := util.DownloadRepo(consts.TemplRepoOwner, consts.TemplRepoName, consts.TemplRepoBranch, zippedRepo)
		if err != nil {
			return err
		}

		// Unzip the repository
		if Verbose {
			fmt.Fprintf(cmd.OutOrStdout(), "Unzipping %s\n", zippedRepo)
		}
		err = util.Unzip(zippedRepo, util.TempPath(""))
		if err != nil {
			return err
		}

		// Sync the templates directory to the working directory
		unzippedRepo := util.TempPath(fmt.Sprintf("%s-%s", consts.TemplRepoName, consts.TemplRepoBranch))
		dst := fmt.Sprintf("%s", consts.DefaultTemplatesDir)
		if Verbose {
			fmt.Fprintf(cmd.OutOrStdout(), "Syncing %s to %s\n", unzippedRepo, dst)
		}
		err = util.SyncDirs(unzippedRepo, dst, Overwrite, cmd.InOrStdin(), cmd.OutOrStdout())
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Successfully pulled template solution files to %s\n", dst)
		return nil
	},
}
