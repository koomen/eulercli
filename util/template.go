package util

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
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
func renderToFile(
	tmpl, dst string,
	data interface{},
	overwrite bool,
	stdin io.Reader,
	stdout io.Writer,
) error {
	// Use the template file to create a template instance
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	stat, err := os.Stat(dst)
	if err == nil {
		if stat.IsDir() {
			return fmt.Errorf("cannot overwrite directory %s", dst)
		}
		if !overwrite {
			confirm, err := Confirm(fmt.Sprintf("Overwrite file %s?", dst), false, stdin, stdout)
			if err != nil {
				return err
			}
			if !confirm {
				return nil
			}
		}
	}

	f, err := CreateFile(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the rendered text into the destination file
	err = t.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

// RenderTemplateDir - Render all files in the tmplDir into dst
func RenderTemplateDir(
	tmplDir,
	dstDir string,
	problem *EulerProblem,
	overwrite bool,
	stdin io.Reader,
	stdout io.Writer,
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

			return renderToFile(tmpl, dst, problem, overwrite, stdin, stdout)
		},
	)
}
