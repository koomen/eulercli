package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingProblemError(t *testing.T) {
	mpe := MissingProblemError{
		ProblemNum:                 700,
		LargestSupportedProblemNum: 500,
	}

	got := mpe.Error()

	want := "Problem 700 not found. The euler CLI supports problems 1-500."

	assert.Equal(t, want, got)
}

func TestGetProblemText(t *testing.T) {
	var problems map[int]string = map[int]string{
		1:   "If we list all the natural numbers",
		469: "In a room N chairs are placed around a round table.",
	}

	for problemNum, substr := range problems {
		got, err := getProblemText(problemNum)
		assert.True(t, strings.Contains(got, substr))
		assert.NoError(t, err)
	}

	// Test error conditition
	_, err := getProblemText(1_000_000)
	assert.Error(t, err)
}

func TestGetAnswer(t *testing.T) {
	var answers map[int]string = map[int]string{
		1:   "233168",
		469: "0.56766764161831",
	}

	for problemNum, want := range answers {
		got, err := getAnswer(problemNum)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	}

	// Test error condition
	_, err := getAnswer(1_000_000)
	assert.Error(t, err)
}

func TestGetProblem(t *testing.T) {
	var eulerProblems map[int]*EulerProblem = map[int]*EulerProblem{
		1: {
			ProblemNum: 1,
			ProblemText: ("If we list all the natural numbers below 10 that are multiples of 3 or 5,\n" +
				"we get 3, 5, 6 and 9. The sum of these multiples is 23.\n" +
				"\n" +
				"Find the sum of all the multiples of 3 or 5 below 1000."),
			AnswerMD5: "e1edf9d1967ca96767dcc2b2d6df69f4",
			Answer:    "233168",
		},
		469: {
			ProblemNum: 469,
			ProblemText: ("In a room N chairs are placed around a round table.\n" +
				"Knights enter the room one by one and choose at random an available empty\n" +
				"chair.\n" +
				"To have enough elbow room the knights always leave at least one empty\n" +
				"chair between each other.\n" +
				"\n" +
				"When there aren't any suitable chairs left, the fraction C of empty chairs\n" +
				"is determined.\n" +
				"We also define E(N) as the expected value of C.\n" +
				"We can verify that E(4) = 1/2 and E(6) = 5/9.\n" +
				"\n" +
				"Find E(10^18). Give your answer rounded to fourteen decimal places in the\n" +
				"form 0.abcdefghijklmn."),
			AnswerMD5: "3c2b641262880db5b735cfa4d4c957bc",
			Answer:    "0.56766764161831",
		},
	}

	for problemNum, want := range eulerProblems {
		got, _ := GetProblem(problemNum)
		assert.Equal(t, want, got)
	}
}

// func TestGetProblemText(t *testing.T) {
// 	testProblems := map[int]string{
// 		1:   "e1edf9d1967ca96767dcc2b2d6df69f4",
// 		469: "3c2b641262880db5b735cfa4d4c957bc",
// 	}

// 	for problem, want := range testProblems {
// 		if got := GetHashedAnswer(problem); got != want {
// 			t.Errorf("GetHashedAnswer(%d) = %s, want %s", problem, got, want)
// 		}
// 	}
// }

// func TestGetHashedAnswer(t *testing.T) {
// 	testProblems := map[int]string{
// 		1: `If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

// Answer: e1edf9d1967ca96767dcc2b2d6df69f4`,
// 		469: `In a room N chairs are placed around a round table.
// Knights enter the room one by one and choose at random an available empty
// chair.
// To have enough elbow room the knights always leave at least one empty
// chair between each other.

// When there aren't any suitable chairs left, the fraction C of empty chairs
// is determined.
// We also define E(N) as the expected value of C.
// We can verify that E(4) = 1/2 and E(6) = 5/9.

// Find E(10^18). Give your answer rounded to fourteen decimal places in the
// form 0.abcdefghijklmn.

// Answer: 3c2b641262880db5b735cfa4d4c957bc`,
// 	}

// 	for problem, want := range testProblems {
// 		if got := GetProblemText(problem); got != want {
// 			t.Errorf("GetProblemText(%d) = %s, want %s", problem, got, want)
// 		}
// 	}
// }
