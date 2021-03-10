package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCmd(t *testing.T) {
	var stdin, stdout bytes.Buffer

	// Check answers via command line params
	rootCmd.SetArgs([]string{"check", "1", "233168"})
	rootCmd.SetOut(&stdout)
	rootCmd.SetIn(&stdin)
	err := rootCmd.Execute()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	want := "Congratulations, 233168 is the correct answer to problem 1!\n"
	assert.Equal(t, want, string(out))

	rootCmd.SetArgs([]string{"check", "1", "233169"})
	rootCmd.SetOut(&stdout)
	rootCmd.SetIn(&stdin)
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err = ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	want = "233169 is not the correct answer to problem 1. Keep trying!\n"
	assert.Equal(t, want, string(out))

	// Check answer via piped input
	rootCmd.SetArgs([]string{"check", "1"})
	stdin.Write([]byte("The answer is 233168.\n"))
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err = ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	want = ("No answer parameter detected. Scanning stdin for correct answer...\n" +
		"-------------------------------------------------------------------------------\n\n\n" +
		"The answer is 233168.\n" +
		"\n\n-------------------------------------------------------------------------------\n" +
		"Detected answer 233168 in input.\n" +
		"Congratulations, this is the correct answer to problem 1!\n")
	assert.Equal(t, want, string(out))

	rootCmd.SetArgs([]string{"check", "1"})
	stdin.Write([]byte("The answer is 233169.\n"))
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err = ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	want = ("No answer parameter detected. Scanning stdin for correct answer...\n" +
		"-------------------------------------------------------------------------------\n\n\n" +
		"The answer is 233169.\n" +
		"\n\n-------------------------------------------------------------------------------\n" +
		"Failed to find correct answer for problem 1 in input. Keep trying!\n")
	assert.Equal(t, want, string(out))

	// Check answer and infer problem number via piped input
	rootCmd.SetArgs([]string{"check"})
	stdin.Write([]byte("The answer to problem 1 is 233168.\n"))
	err = rootCmd.Execute()
	assert.NoError(t, err)
	out, err = ioutil.ReadAll(&stdout)
	assert.NoError(t, err)
	want = ("No parameters detected. Scanning stdin for problem number and correct answer...\n" +
		"-------------------------------------------------------------------------------\n\n\n" +
		"The answer to problem 1 is 233168.\n" +
		"\n\n-------------------------------------------------------------------------------\n" +
		"Extracted problem number 1 from input\n" +
		"Detected answer 233168 in input.\n" +
		"Congratulations, this is the correct answer to problem 1!\n")
	assert.Equal(t, want, string(out))
}
