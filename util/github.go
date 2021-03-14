package util

import (
	"fmt"
)

// BuildRepoUrl - construct a URL for the given github repo
func BuildRepoUrl(owner, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s", owner, repo)
}

// BuildRepoArchiveURL - construct a URL pointing to a zipped Github Repo
func BuildRepoArchiveUrl(owner, repo, branch string) string {
	return fmt.Sprintf("https://github.com/%s/%s/archive/%s.zip", owner, repo, branch)
}

// DownloadRepo - download a zipped copy of the specified GitHub repo
func DownloadRepo(owner, repo, branch, dst string) error {
	url := BuildRepoArchiveUrl(owner, repo, branch)
	return DownloadFile(url, dst)
}
