package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/koomen/eulercli/util"
	"github.com/stretchr/testify/assert"
)

func TestAnswerCmd(t *testing.T) {
	var stdout bytes.Buffer

	rootCmd.SetArgs([]string{"answer", "1"})
	rootCmd.SetOut(&stdout)
	err := rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	assert.Equal(t, "e1edf9d1967ca96767dcc2b2d6df69f4\n", string(out))
}

func TestAnswerCmdProblemOutOfBounds(t *testing.T) {
	var stdout bytes.Buffer

	rootCmd.SetArgs([]string{"answer", "1000000"})
	rootCmd.SetOut(&stdout)
	err := rootCmd.Execute()
	_, ok := err.(*util.MissingProblemError)
	assert.True(t, ok)
}
