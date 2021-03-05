package util

import (
	"fmt"
)

// BuildRepoArchiveURL - construct a URL pointing to a zipped Github Repo
func BuildRepoArchiveURL(owner, repo, branch string) string {
	return fmt.Sprintf("https://github.com/%s/%s/archive/%s.zip", owner, repo, branch)
}

// DownloadRepo - download a zipped copy of the specified GitHub repo
func DownloadRepo(owner, repo, branch, dst string) error {
	url := BuildRepoArchiveURL(owner, repo, branch)
	return DownloadFile(url, dst)
}
