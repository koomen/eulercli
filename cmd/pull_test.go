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

	_, err := os.Stat("eulercli_templates")
	assert.True(t, os.IsNotExist(err))

	rootCmd.SetArgs([]string{"pull"})
	rootCmd.SetOut(&stdout)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)

	wants := []string{
		"Downloading templates from https://github.com/koomen/eulercli\n",
		"Successfully pulled template solution files to eulercli_templates\n",
	}
	for _, want := range wants {
		assert.Contains(t, string(out), want)
	}

	assert.DirExists(t, "eulercli_templates")
	assert.DirExists(t, "eulercli_templates/julia")
	assert.FileExists(t, "eulercli_templates/julia/initenv.jl")
	assert.FileExists(t, "eulercli_templates/julia/src/euler{{.PaddedProblemNum}}/solution.jl")

	os.RemoveAll("eulercli_templates")
}

func TestPullCmdVerbose(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var stdout bytes.Buffer

	_, err := os.Stat("eulercli_templates")
	assert.True(t, os.IsNotExist(err))

	rootCmd.SetArgs([]string{"pull", "-v"})
	rootCmd.SetOut(&stdout)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)

	wants := []string{
		"Downloading templates from https://github.com/koomen/eulercli\n",
		"Unzipping /tmp/eulercli/eulercli-main.zip\n",
		"Syncing /tmp/eulercli/eulercli-main/templates to eulercli_templates\n",
		"Successfully pulled template solution files to eulercli_templates\n",
	}
	for _, want := range wants {
		assert.Contains(t, string(out), want)
	}

	assert.DirExists(t, "eulercli_templates")
	assert.DirExists(t, "eulercli_templates/julia")
	assert.FileExists(t, "eulercli_templates/julia/initenv.jl")
	assert.FileExists(t, "eulercli_templates/julia/src/euler{{.PaddedProblemNum}}/solution.jl")

	os.RemoveAll("eulercli_templates")
}
