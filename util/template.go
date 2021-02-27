package util

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/koomen/eulercli/consts"
)

// renderTemplateToString - Takes a template string and data instance and returns a
// rendered string
func renderTemplateToString(templStr string, data interface{}) (string, error) {
	t, err := template.New("t").Parse(templStr)
	if err != nil {
		return "", err
	}

	var res bytes.Buffer
	if execErr := t.Execute(&res, data); execErr != nil {
		return "", execErr
	}

	return res.String(), nil
}

// createFile - Creates a file and its directory if needed
func createFile(path string) (*os.File, error) {
	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), consts.DirPerm); err != nil {
		return nil, err
	}

	return os.Create(path)
}

// renderTemplateToFile - Takes a template file and data instance to renders to a
// destination file
func renderTemplateToFile(templatePath string, dest io.Writer, data interface{}) error {
	// Use the template file to create a template instance
	t, templErr := template.ParseFiles(templatePath)
	if templErr != nil {
		return templErr
	}

	// Write the rendered text into the destination file
	return t.Execute(dest, data)
}

// RenderFiles - Render all files in the templatePath into the outputPath
func RenderFiles(
	templateDirPath,
	destDirPath string,
	problem *EulerProblem,
	overwriteExisting bool,
) error {
	return filepath.Walk(templateDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories
			if info.IsDir() {
				return nil
			}

			// build the "raw" output path by appending the relative path
			// to the output path
			relPath, relErr := filepath.Rel(templateDirPath, path)
			if relErr != nil {
				return relErr
			}
			rawDestPath := filepath.Join(destDirPath, relPath)

			// Render any template variables in the output path
			// e.g. euler{{.ProblemNum}}.jl -> euler25.jl
			destPath, renderErr := renderTemplateToString(rawDestPath, problem)
			if renderErr != nil {
				return renderErr
			}

			// Only write the template file if the destination file does not
			// already exist or the overwriteExisting parameter is true
			if _, err := os.Stat(destPath); os.IsNotExist(err) || overwriteExisting {
				fmt.Printf("Writing template %s to %s\n", path, destPath)
				dest, fileErr := createFile(destPath)
				if fileErr != nil {
					return fileErr
				}
				defer dest.Close()
				return renderTemplateToFile(path, dest, problem)
			} else {
				fmt.Printf("Skipped existing file %s. Use --overwrite to overwrite existing files.\n", destPath)
				return nil
			}
		},
	)
}
