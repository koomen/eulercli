package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var problem = &EulerProblem{
	ProblemNum:  25,
	ProblemText: "Here is the problem text",
	AnswerMD5:   "theanswer",
}

const templateStr = "Problem {{.ProblemNum}} answer is {{.AnswerMD5}}"
const badTemplateStr = "Problem {{.NotAField}} answer is {{.AnswerMD5}}"
const rendered = "Problem 25 answer is theanswer"

func TestRenderTemplateToString(t *testing.T) {
	want := "Problem 25 answer is theanswer"
	got, err := renderTemplateToString(templateStr, problem)

	assert.NoError(t, err)
	assert.Equal(t, want, got)

	_, err2 := renderTemplateToString(badTemplateStr, problem)
	assert.Error(t, err2)
}

func TestRenderTemplateToFile(t *testing.T) {
	templatePath := "/tmp/template.txt"
	assert.NoError(t, ioutil.WriteFile(templatePath, []byte(templateStr), 0766))
	defer os.Remove(templatePath)

	var dest bytes.Buffer
	assert.NoError(t, renderTemplateToFile(templatePath, &dest, problem))

	want, got := rendered, dest.String()
	assert.Equal(t, want, got)

	assert.NoError(t, ioutil.WriteFile(templatePath, []byte(badTemplateStr), 0766))
	assert.Error(t, renderTemplateToFile(templatePath, &dest, problem))
}

func TestRenderFiles(t *testing.T) {
	templateDir := "/tmp/template"
	os.RemoveAll(templateDir)
	var templates map[string]string = map[string]string{
		"file1.txt": templateStr,
		"dir{{.ProblemNum}}/file{{.AnswerMD5}}.txt": "Template text",
	}

	for path, s := range templates {
		f, err := createFile(filepath.Join(templateDir, path))
		assert.NoError(t, err)
		_, writeErr := f.Write([]byte(s))
		assert.NoError(t, writeErr)
		assert.NoError(t, f.Close())
	}

	destDir := "/tmp/output"
	os.RemoveAll(destDir)
	var expected map[string]string = map[string]string{
		"file1.txt":               rendered,
		"dir25/filetheanswer.txt": "Template text",
	}

	assert.NoError(t, RenderFiles(templateDir, destDir, problem, true))

	for path, want := range expected {
		got, err := ioutil.ReadFile(filepath.Join(destDir, path))
		assert.NoError(t, err)
		assert.Equal(t, want, string(got))
	}

	os.RemoveAll(templateDir)
	os.RemoveAll(destDir)
}
