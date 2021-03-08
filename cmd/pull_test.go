package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullCmd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var stdout bytes.Buffer

	_, err := os.Stat("./eulercli_templates")
	assert.True(t, os.IsNotExist(err))

	rootCmd.SetArgs([]string{"pull"})
	rootCmd.SetOut(&stdout)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)

	want := ("Downloading templates from https://github.com/koomen/eulercli\n" +
		"Successfully pulled template solution files to ./eulercli_templates\n")
	assert.Equal(t, want, string(out))

	assert.DirExists(t, "./eulercli_templates")
	assert.DirExists(t, "./eulercli_templates/julia")
	assert.FileExists(t, "./eulercli_templates/julia/template.jl")

	os.RemoveAll("./eulercli_templates")
}

func TestPullCmdVerbose(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var stdout bytes.Buffer

	_, err := os.Stat("./eulercli_templates")
	assert.True(t, os.IsNotExist(err))

	rootCmd.SetArgs([]string{"pull", "-v"})
	rootCmd.SetOut(&stdout)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)

	want := ("Downloading templates from https://github.com/koomen/eulercli\n" +
		"Unzipping /tmp/eulercli/eulercli-main.zip\n" +
		"Syncing /tmp/eulercli/eulercli-main/templates to ./eulercli_templates\n" +
		"Successfully pulled template solution files to ./eulercli_templates\n")
	assert.Equal(t, want, string(out))

	assert.DirExists(t, "./eulercli_templates")
	assert.DirExists(t, "./eulercli_templates/julia")
	assert.FileExists(t, "./eulercli_templates/julia/template.jl")

	os.RemoveAll("./eulercli_templates")
}
