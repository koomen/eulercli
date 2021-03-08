package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/koomen/eulercli/util"
	"github.com/stretchr/testify/assert"
)

func TestProblemCmd(t *testing.T) {
	var stdout bytes.Buffer

	rootCmd.SetArgs([]string{"problem", "1"})
	rootCmd.SetOut(&stdout)
	err := rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	want := ("If we list all the natural numbers below 10 that are multiples of 3 or 5,\n" +
		"we get 3, 5, 6 and 9. The sum of these multiples is 23.\n" +
		"\n" +
		"Find the sum of all the multiples of 3 or 5 below 1000.\n")
	assert.Equal(t, want, string(out))
}

func TestProblemCmdProblemOutOfBounds(t *testing.T) {
	var stdout bytes.Buffer

	rootCmd.SetArgs([]string{"problem", "1000000"})
	rootCmd.SetOut(&stdout)
	err := rootCmd.Execute()
	_, ok := err.(*util.MissingProblemError)
	assert.True(t, ok)
}
