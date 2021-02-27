package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemArgError(t *testing.T) {
	pae := ProblemArgError{"hello"}
	got := pae.Error()
	want := "Problem arg \"hello\" is not an integer"
	assert.Equal(t, want, got)
}

func TestValidateProblemStr(t *testing.T) {
	err := ValidateProblemStr("123")
	assert.NoError(t, err)
	err = ValidateProblemStr("hello")
	assert.Error(t, err)
	got := err.Error()
	want := "Problem arg \"hello\" is not an integer"
	assert.Equal(t, want, got)
}

func TestValidateSingleProblemArg(t *testing.T) {
	err := ValidateSingleProblemArg(nil, []string{"42"})
	assert.NoError(t, err)
	err = ValidateSingleProblemArg(nil, []string{"hello"})
	assert.Error(t, err)
	err = ValidateSingleProblemArg(nil, []string{})
	assert.Error(t, err)
}
