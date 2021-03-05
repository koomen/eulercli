package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRepoArchiveURL(t *testing.T) {
	want := "https://github.com/koomen/eulercli/archive/main.zip"
	got := BuildRepoArchiveURL("koomen", "eulercli", "main")

	assert.Equal(t, want, got)
}

func TestDownloadRepo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	dst := "/tmp/test1.zip"

	os.Remove(dst)

	err := DownloadRepo("koomen", "eulercli", "main", dst)
	assert.NoError(t, err)
	defer os.Remove(dst)

	_, err = os.Stat(dst)
	assert.NoError(t, err)
}
