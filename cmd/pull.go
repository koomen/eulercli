package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/koomen/eulercli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

// pullCmd
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Download solution templates to ./templates",
	Long:  fmt.Sprintf("Download solution templates to ./templates"),
	RunE: func(cmd *cobra.Command, args []string) error {
		util.CreateTempDir()
		defer util.RemoveTempDir()

		owner, repo, branch := "koomen", "eulercli", "main"

		fmt.Fprintf(cmd.OutOrStdout(), "Downloading templates from https://github.com/%s/%s\n", owner, repo)

		zippedRepo := util.TempPath(fmt.Sprintf("%s-%s.zip", repo, branch))
		err := util.DownloadRepo(owner, repo, branch, zippedRepo)
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
		unzippedRepo := util.TempPath(fmt.Sprintf("%s-%s", repo, branch))
		tmplDir := filepath.Join(unzippedRepo, "templates")
		dst := "./templates"
		if Verbose {
			fmt.Fprintf(cmd.OutOrStdout(), "Syncing %s to %s\n", tmplDir, dst)
		}
		err = util.SyncDirs(tmplDir, dst, false, os.Stdin)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Successfully pulled template solution files to ./templates\n")
		return nil
	},
}
