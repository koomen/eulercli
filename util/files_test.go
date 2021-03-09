package util

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/koomen/eulercli/consts"
	"github.com/stretchr/testify/assert"
)

func TestTempDir(t *testing.T) {
	// Ensure temp directory does not exist
	err := RemoveTempDir()
	assert.NoError(t, err)

	_, err = os.Stat(consts.TempDirPath)
	assert.True(t, os.IsNotExist(err))

	// Create temp dir and verify it exists
	err = CreateTempDir()
	assert.NoError(t, err)
	_, err = os.Stat(consts.TempDirPath)
	assert.False(t, os.IsNotExist(err))

	// Put a file into the newly created temp dir
	path := TempPath("file.txt")
	err = os.WriteFile(path, []byte(""), consts.FilePerm)
	assert.NoError(t, err)

	// Ensure file exists
	_, err = os.Stat(path)
	assert.False(t, os.IsNotExist(err))

	// Recreate the temp dir
	assert.NoError(t, CreateTempDir())
	_, err = os.Stat(consts.TempDirPath)
	assert.False(t, os.IsNotExist(err))

	// Ensure file no longer exists
	_, err = os.Stat(path)
	assert.True(t, os.IsNotExist(err))

	// Remove the temp dir
	RemoveTempDir()
	_, err = os.Stat(consts.TempDirPath)
	assert.True(t, os.IsNotExist(err))

}

func TestCreateFile(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	f, err := CreateFile(TempPath("dir1/file1.txt"))
	assert.NoError(t, err)

	_, err = f.WriteString("hello\n")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	want := "hello\n"
	got, err := os.ReadFile(TempPath("dir1/file1.txt"))
	assert.NoError(t, err)
	assert.Equal(t, want, string(got))

}

func TestDownloadFile(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	dl1, dl2 := TempPath("dl1"), TempPath("dl2")

	err := DownloadFile("https://example.com", dl1)
	assert.NoError(t, err)
	_, err = os.Stat(dl1)
	assert.False(t, os.IsNotExist(err))

	err = DownloadFile("https://doesnotexist.example.com", dl2)
	assert.Error(t, err)
	_, err = os.Stat(dl2)
	assert.True(t, os.IsNotExist(err))
}

func TestConfirm(t *testing.T) {
	var stdin, stdout bytes.Buffer
	msg := "Enter confirmation"

	// Test "yes" response
	stdin.WriteString("y\n")
	wantResult := true
	wantOut := fmt.Sprintf("%s [Yn]:", msg)

	gotResult, err := Confirm(msg, true, &stdin, &stdout)
	assert.NoError(t, err)

	gotOut, err := stdout.ReadString(':')
	assert.NoError(t, err)
	assert.Equal(t, wantResult, gotResult)
	assert.Equal(t, wantOut, gotOut)

	// Test "no" response
	stdin.Reset()
	stdout.Reset()

	stdin.WriteString("n\n")
	wantResult = false
	wantOut = fmt.Sprintf("%s [Yn]:", msg)

	gotResult, err = Confirm(msg, true, &stdin, &stdout)
	assert.NoError(t, err)

	gotOut, err = stdout.ReadString(':')
	assert.NoError(t, err)
	assert.Equal(t, wantResult, gotResult)
	assert.Equal(t, wantOut, gotOut)

	// Test default response
	stdin.Reset()
	stdout.Reset()

	stdin.WriteString("\n")
	wantResult = false
	wantOut = fmt.Sprintf("%s [yN]:", msg)

	gotResult, err = Confirm(msg, false, &stdin, &stdout)
	assert.NoError(t, err)

	gotOut, err = stdout.ReadString(':')
	assert.NoError(t, err)
	assert.Equal(t, wantResult, gotResult)
	assert.Equal(t, wantOut, gotOut)

	// Test unrecognized response
	stdin.Reset()
	stdout.Reset()

	stdin.WriteString("yes\nn\n")
	wantResult = false
	wantOut = fmt.Sprintf("%s [Yn]: Response \"yes\" not recognized.\n", msg)

	gotResult, err = Confirm(msg, true, &stdin, &stdout)
	assert.NoError(t, err)

	gotOut, err = stdout.ReadString('\n')
	assert.NoError(t, err)
	assert.Equal(t, wantResult, gotResult)
	assert.Equal(t, wantOut, gotOut)

}

func TestComputeFileChecksum(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	path := TempPath("file.txt")
	err := os.WriteFile(path, []byte("hello\n"), consts.FilePerm)
	assert.NoError(t, err)

	var want uint32 = 0x363a3020
	got, err := ComputeFileChecksum(path)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestAreChecksumsEqual(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	// Create two identical files
	p1, p2 := TempPath("file1.txt"), TempPath("file2.txt")
	err := os.WriteFile(p1, []byte("hello\n"), consts.FilePerm)
	assert.NoError(t, err)
	err = os.WriteFile(p2, []byte("hello\n"), consts.FilePerm)
	assert.NoError(t, err)

	// Checksums should be equal
	got, err := AreChecksumsEqual(p1, p2)
	assert.True(t, got)

	// Create a new non-identical file
	p3 := TempPath("file3.txt")
	err = os.WriteFile(p3, []byte("goodbye\n"), consts.FilePerm)
	assert.NoError(t, err)

	// Checksums should be different
	got, err = AreChecksumsEqual(p1, p3)
	assert.False(t, got)
}

func TestUnzip(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	Unzip("test/dir1.zip", TempPath(""))

	assert.FileExists(t, TempPath("dir1/foo.txt"))
	assert.FileExists(t, TempPath("dir1/dir2/bar.txt"))
}

func TestSyncFiles(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	var stdin bytes.Buffer

	f1 := TempPath("dir1/file1.txt")
	f2 := TempPath("dir1/dir2/file2.txt")

	// Create the following directory structure
	// 		dir1/
	//		dir1/file1.txt
	//		dir1/dir2
	//		dir1/dir2/file2.txt
	os.MkdirAll(TempPath("dir1/dir2"), consts.DirPerm)
	os.WriteFile(f1, []byte("hello\n"), consts.FilePerm)
	os.WriteFile(f2, []byte("goodbye\n"), consts.FilePerm)
	_, err := os.Stat(f1)
	assert.False(t, os.IsNotExist(err))
	_, err = os.Stat(f2)
	assert.False(t, os.IsNotExist(err))

	// Sync files
	err = SyncFiles(f1, TempPath("destdir/file1.txt"), false, &stdin)
	assert.NoError(t, err)
	err = SyncFiles(f2, TempPath("destdir/file2.txt"), false, &stdin)
	assert.NoError(t, err)

	// Ensure sync'd files are equal
	got, err := AreChecksumsEqual(f1, TempPath("destdir/file1.txt"))
	assert.True(t, got)
	got, err = AreChecksumsEqual(f2, TempPath("destdir/file2.txt"))
	assert.True(t, got)

	// Sync f2 -> distdir/file1.txt (don't overwrite)
	fmt.Println("ready")
	stdin.WriteString("n\n")
	err = SyncFiles(f2, TempPath("destdir/file1.txt"), false, &stdin)
	got, err = AreChecksumsEqual(f2, TempPath("destdir/file1.txt"))
	assert.False(t, got)

	// Sync f2 -> distdir/file1.txt (overwrite)
	stdin.Reset()
	stdin.WriteString("y\n")
	err = SyncFiles(f2, TempPath("destdir/file1.txt"), false, &stdin)
	got, err = AreChecksumsEqual(f2, TempPath("destdir/file1.txt"))
	assert.True(t, got)

	// Sync f1 -> distdir/file1.txt (overwrite)
	stdin.Reset()
	stdin.WriteString("n\n")
	err = SyncFiles(f1, TempPath("destdir/file1.txt"), true, &stdin)
	got, err = AreChecksumsEqual(f1, TempPath("destdir/file1.txt"))
	assert.True(t, got)
}

func TestSyncDirs(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	var stdin bytes.Buffer

	d1 := TempPath("dir1")
	// d2 := TempPath("dir1/dir2")
	d3 := TempPath("dir3")
	f1 := TempPath("dir1/file1.txt")
	f2 := TempPath("dir1/dir2/file2.txt")

	// Create the following directory structure
	// 		dir1/
	//		dir1/file1.txt
	//		dir1/dir2
	//		dir1/dir2/file2.txt
	os.MkdirAll(TempPath("dir1/dir2"), consts.DirPerm)
	os.WriteFile(f1, []byte("hello\n"), consts.FilePerm)
	os.WriteFile(f2, []byte("goodbye\n"), consts.FilePerm)
	_, err := os.Stat(f1)
	assert.False(t, os.IsNotExist(err))
	_, err = os.Stat(f2)
	assert.False(t, os.IsNotExist(err))

	// Sync dir1 to dir3
	err = SyncDirs(d1, d3, false, &stdin)
	assert.NoError(t, err)

	// source and destination files should be equal
	got, err := AreChecksumsEqual(f1, filepath.Join(d3, "file1.txt"))
	assert.NoError(t, err)
	assert.True(t, got)
	got, err = AreChecksumsEqual(f2, filepath.Join(d3, "dir2/file2.txt"))
	assert.NoError(t, err)
	assert.True(t, got)

	// Delete dir3/dir2/file2.txt, and change the contents of dir1/file1.txt
	os.Remove(filepath.Join(d3, "dir2/file2.txt"))
	os.WriteFile(f1, []byte("hello2\n"), consts.FilePerm)

	// Sync dir1 to dir3
	stdin.Reset()
	stdin.WriteString("n\n")
	err = SyncDirs(d1, d3, false, &stdin)
	assert.NoError(t, err)

	// This time, dir1/file1.txt and dir3/file1.txt should be different
	got, err = AreChecksumsEqual(f1, filepath.Join(d3, "file1.txt"))
	assert.NoError(t, err)
	assert.False(t, got)
	got, err = AreChecksumsEqual(f2, filepath.Join(d3, "dir2/file2.txt"))
	assert.NoError(t, err)
	assert.True(t, got)

	// Sync dir1 to dir3 again, but confirm the overwrite of file1.txt
	stdin.Reset()
	stdin.WriteString("y\n")
	err = SyncDirs(d1, d3, false, &stdin)
	assert.NoError(t, err)

	// This time, dir1/file1.txt and dir3/file1.txt should be identical
	got, err = AreChecksumsEqual(f1, filepath.Join(d3, "file1.txt"))
	assert.NoError(t, err)
	assert.True(t, got)
	got, err = AreChecksumsEqual(f2, filepath.Join(d3, "dir2/file2.txt"))
	assert.NoError(t, err)
	assert.True(t, got)

}
