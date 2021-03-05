package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
		// defer util.RemoveTempDir()

		owner, repo, branch := "koomen", "eulercli", "main"

		zippedRepo := util.TempPath(fmt.Sprintf("%s-%s.zip", repo, branch))
		err := util.DownloadRepo(owner, repo, branch, zippedRepo)
		cobra.CheckErr(err)

		// Unzip the repository
		unzipCmd := exec.Command("unzip", "-o", zippedRepo, "-d", util.TempPath(""))
		//unzipCmd.Stdout = os.Stdout
		err = unzipCmd.Run()
		cobra.CheckErr(err)

		unzippedRepo := util.TempPath(fmt.Sprintf("%s-%s", repo, branch))
		tmplDir := filepath.Join(unzippedRepo, "templates")

		err = util.SyncDirs(tmplDir, "./templates", false, os.Stdin)
		cobra.CheckErr(err)
	},
}
