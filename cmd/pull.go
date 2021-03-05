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
	Run: func(cmd *cobra.Command, args []string) {
		util.CreateTempDir()
		defer util.RemoveTempDir()

		owner, repo, branch := "koomen", "eulercli", "main"

		fmt.Printf("Downloading templates from https://github.com/%s/%s\n", owner, repo)

		zippedRepo := util.TempPath(fmt.Sprintf("%s-%s.zip", repo, branch))
		err := util.DownloadRepo(owner, repo, branch, zippedRepo)
		cobra.CheckErr(err)

		// Unzip the repository
		if Verbose {
			fmt.Printf("Unzipping %s\n", zippedRepo)
		}
		err = util.Unzip(zippedRepo, util.TempPath(""))
		cobra.CheckErr(err)

		// Sync the templates directory to the working directory
		unzippedRepo := util.TempPath(fmt.Sprintf("%s-%s", repo, branch))
		tmplDir := filepath.Join(unzippedRepo, "templates")
		dst := "./templates"
		if Verbose {
			fmt.Printf("Syncing %s to ./templates\n", tmplDir)
		}
		err = util.SyncDirs(tmplDir, dst, false, os.Stdin)
		cobra.CheckErr(err)
		fmt.Printf("Successfully pulled template solution files to ./templates\n")
	},
}
