package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCmd(t *testing.T) {
	var stdin, stdout bytes.Buffer

	// Assert that templates and rendered files don't exists
	_, err := os.Stat("eulercli_templates")
	assert.True(t, os.IsNotExist(err))
	_, err = os.Stat("julia")
	assert.True(t, os.IsNotExist(err))

	// Execute the generate command
	rootCmd.SetArgs([]string{"generate", "1", "--language", "julia"})
	rootCmd.SetOut(&stdout)
	rootCmd.SetIn(&stdin)
	stdin.WriteString("y\n") // confirm the pull prompt
	err = rootCmd.Execute()
	assert.NoError(t, err)

	// Assert that templates and rendered files *do* exist
	assert.DirExists(t, "eulercli_templates")
	assert.FileExists(t, "julia/initenv.jl")
	assert.FileExists(t, "julia/euler0001/solution.jl")
	defer os.RemoveAll("eulercli_templates")
	defer os.RemoveAll("julia")

	rootCmd.SetArgs([]string{"generate", "2", "--language", "julia"})
	err = rootCmd.Execute()
	assert.NoError(t, err)

	assert.FileExists(t, "julia/euler0002/solution.jl")
}
