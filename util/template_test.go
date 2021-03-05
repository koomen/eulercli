package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/koomen/eulercli/consts"
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

func TestRenderToString(t *testing.T) {
	want := "Problem 25 answer is theanswer"
	got, err := renderToString(templateStr, problem)

	assert.NoError(t, err)
	assert.Equal(t, want, got)

	_, err2 := renderToString(badTemplateStr, problem)
	assert.Error(t, err2)
}

func TestRenderToFile(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	// Create template file
	tmpl := TempPath("template.txt")
	assert.NoError(t, ioutil.WriteFile(tmpl, []byte(templateStr), 0766))

	// Render template file to dst.txt
	dst := TempPath("dst.txt")
	assert.NoError(t, renderToFile(tmpl, dst, problem))

	// Ensure dst.txt contains the expected text
	want := rendered
	got, err := os.ReadFile(dst)
	assert.NoError(t, err)
	assert.Equal(t, want, string(got))

	// Write a broken template string into template.txt
	assert.NoError(t, ioutil.WriteFile(tmpl, []byte(badTemplateStr), 0766))

	// renderToFile should error
	assert.Error(t, renderToFile(tmpl, dst, problem))
}

func TestRenderFiles(t *testing.T) {
	CreateTempDir()
	defer RemoveTempDir()

	// Create the following files
	// 		file1.txt
	//		dir{{.ProblemNum}}/file{{.AnswerMD5}}.txt
	os.MkdirAll(TempPath("templates/dir{{.ProblemNum}}"), consts.DirPerm)

	f1 := TempPath("templates/file1.txt")
	err := os.WriteFile(f1, []byte(templateStr), consts.FilePerm)
	assert.NoError(t, err)

	f2 := TempPath("templates/dir{{.ProblemNum}}/file{{.AnswerMD5}}.txt")
	err = os.WriteFile(f2, []byte("Template text"), consts.FilePerm)
	assert.NoError(t, err)

	// Render templates
	dstDir := TempPath("output")
	assert.NoError(t, RenderTemplateDir(TempPath("templates"), dstDir, problem, true))

	want := rendered
	got, err := os.ReadFile(filepath.Join(dstDir, "file1.txt"))
	assert.NoError(t, err)
	assert.Equal(t, want, string(got))

	want = "Template text"
	got, err = os.ReadFile(filepath.Join(dstDir, "dir25/filetheanswer.txt"))
	assert.NoError(t, err)
	assert.Equal(t, want, string(got))
}
