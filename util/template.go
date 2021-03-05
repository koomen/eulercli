package util

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// renderToString - Takes a template string and data instance and returns a
// rendered string
func renderToString(templStr string, data interface{}) (string, error) {
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

// renderToFile - Takes a template file and data instance to renders to a
// destination file
func renderToFile(tmpl, dst string, data interface{}) error {
	// Use the template file to create a template instance
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	f, err := CreateFile(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the rendered text into the destination file
	return t.Execute(f, data)
}

// RenderTemplateDir - Render all files in the tmplDir into dst
func RenderTemplateDir(
	tmplDir,
	dstDir string,
	problem *EulerProblem,
	overwrite bool,
) error {
	return filepath.Walk(tmplDir,
		func(tmpl string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories
			if info.IsDir() {
				return nil
			}

			// build the "raw" output path by appending the relative path
			// to the output path
			relPath, err := filepath.Rel(tmplDir, tmpl)
			if err != nil {
				return err
			}
			rawDst := filepath.Join(dstDir, relPath)

			// Render any template variables in the output path
			// e.g. euler{{.ProblemNum}}.jl -> euler25.jl
			dst, err := renderToString(rawDst, problem)
			if err != nil {
				return err
			}

			// Only write the template file if the destination file does not
			// already exist or the overwriteExisting parameter is true
			_, err = os.Stat(dst)
			if os.IsNotExist(err) ||
				overwrite ||
				Confirm(fmt.Sprintf("Overwrite file %s", dst), false, os.Stdin, os.Stdout) {
				fmt.Printf("Writing template %s to %s\n", tmpl, dst)
				return renderToFile(tmpl, dst, problem)
			}

			return nil
		},
	)
}
