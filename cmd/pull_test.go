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

	_, err := os.Stat("./templates")
	assert.True(t, os.IsNotExist(err))

	rootCmd.SetArgs([]string{"pull"})
	rootCmd.SetOut(&stdout)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)

	want := ("Downloading templates from https://github.com/koomen/eulercli\n" +
		"Successfully pulled template solution files to ./templates\n")
	assert.Equal(t, want, string(out))

	assert.DirExists(t, "./templates")
	assert.DirExists(t, "./templates/julia")
	assert.FileExists(t, "./templates/julia/template.jl")

	os.RemoveAll("./templates")
}

func TestPullCmdVerbose(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var stdout bytes.Buffer

	_, err := os.Stat("./templates")
	assert.True(t, os.IsNotExist(err))

	rootCmd.SetArgs([]string{"pull", "-v"})
	rootCmd.SetOut(&stdout)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)

	want := ("Downloading templates from https://github.com/koomen/eulercli\n" +
		"Unzipping /tmp/eulercli/eulercli-main.zip\n" +
		"Syncing /tmp/eulercli/eulercli-main/templates to ./templates\n" +
		"Successfully pulled template solution files to ./templates\n")
	assert.Equal(t, want, string(out))

	assert.DirExists(t, "./templates")
	assert.DirExists(t, "./templates/julia")
	assert.FileExists(t, "./templates/julia/template.jl")

	os.RemoveAll("./templates")
}
