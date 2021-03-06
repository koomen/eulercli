package util

import (
	"archive/zip"
	"bufio"
	"fmt"
	"hash/crc32"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/koomen/eulercli/consts"
)

// RemoveTempDir - Remove the temporary directory
func RemoveTempDir() error {
	return os.RemoveAll(consts.TempDirPath)
}

// CreateTempDir - Create a temporary directory
func CreateTempDir() error {
	RemoveTempDir()
	return os.MkdirAll(consts.TempDirPath, consts.DirPerm)
}

// CreateFile - Creates a file and its directory if needed
func CreateFile(path string) (*os.File, error) {
	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), consts.DirPerm); err != nil {
		return nil, err
	}

	return os.Create(path)
}

// TempPath - Append a relative path onto eulercli's temporary directory
func TempPath(relPath string) string {
	return filepath.Join(consts.TempDirPath, relPath)
}

// DownloadFile - download a file and save it to the specified destination
func DownloadFile(url, dst string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := CreateFile(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// Confirm - get a confirmation from the user. A "y" response will return true
// and an "n" response will return false
func Confirm(msg string, defaultVal bool, stdin io.Reader, stdout io.Writer) (bool, error) {
	var options string
	if defaultVal {
		options = "[Yn]"
	} else {
		options = "[yN]"
	}

	reader := bufio.NewReader(stdin)
	for true {
		stdout.Write([]byte(fmt.Sprintf("%s %s: ", msg, options)))
		resp, _ := reader.ReadString('\n')
		if len(resp) < 1 {
			return false, fmt.Errorf("failed to read from stdin")
		}
		trimmed := resp[0 : len(resp)-1]
		switch trimmed {
		case "":
			return defaultVal, nil
		case "y", "Y":
			return true, nil
		case "n", "N":
			return false, nil
		}
		stdout.Write([]byte(fmt.Sprintf("Response \"%s\" not recognized.\n", trimmed)))
	}

	// Unreachable statement; included for the compiler
	return false, nil
}

// ComputeFileChecksum - compute the crc32 checksum of a file
func ComputeFileChecksum(file string) (uint32, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}
	return crc32.ChecksumIEEE(dat), nil
}

// AreChecksumsEqual - compare checksums for two files
func AreChecksumsEqual(f1, f2 string) (bool, error) {
	cs1, err := ComputeFileChecksum(f1)
	if err != nil {
		return false, err
	}
	cs2, err := ComputeFileChecksum(f2)
	if err != nil {
		return false, err
	}
	return cs1 == cs2, nil
}

// Unzip - Unzip a zip archive
// Taken from https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func Unzip(src, dst string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dst, consts.DirPerm)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dst, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dst)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

// SyncFiles - Sync a source file a destination
// If the destination exists and overwrite is false, the user is asked to confirm
// the operation.
func SyncFiles(src, dst string, overwrite bool, stdin io.Reader, stdout io.Writer) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if srcStat.IsDir() {
		return fmt.Errorf("SyncFiles source %s is a directory", src)
	}
	dstStat, err := os.Stat(dst)
	if err == nil {
		if dstStat.IsDir() {
			return fmt.Errorf("SyncFiles destination %s is a directory", dst)
		}
		areEqual, err := AreChecksumsEqual(src, dst)
		if err != nil {
			return err
		}
		if areEqual {
			// Todo: src and dst are equal; should return (and handle) an error
			return nil
		}
		if !overwrite {
			msg := fmt.Sprintf("File %s exists. Overwrite?", dst)
			confirm, err := Confirm(msg, false, stdin, stdout)
			if err != nil {
				return err
			}
			if !confirm {
				// Todo: the user did not confirm; should return (and handle) an error
				return nil
			}
		}
	}

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dst), consts.DirPerm); err != nil {
		return err
	}

	srcDat, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, srcDat, consts.FilePerm)
	if err != nil {
		return err
	}
	stdout.Write([]byte(fmt.Sprintf("  Wrote file %s\n", dst)))
	return nil
}

// SyncDirs - Sync a source directory to a destination
func SyncDirs(src, dst string, overwrite bool, stdin io.Reader, stdout io.Writer) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcStat.IsDir() {
		return fmt.Errorf("SyncDirs source %s is not a directory", src)
	}

	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)
		return SyncFiles(path, dstPath, overwrite, stdin, stdout)
	})
}
