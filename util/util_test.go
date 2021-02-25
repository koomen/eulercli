package util

import (
	"testing"
)

func TestGetProblemText(t *testing.T) {
	testProblems := map[int]string{
		1:   "e1edf9d1967ca96767dcc2b2d6df69f4",
		469: "3c2b641262880db5b735cfa4d4c957bc",
	}

	for problem, want := range testProblems {
		if got := GetHashedAnswer(problem); got != want {
			t.Errorf("GetHashedAnswer(%d) = %s, want %s", problem, got, want)
		}
	}
}

func TestGetHashedAnswer(t *testing.T) {
	testProblems := map[int]string{
		1: `If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.


Answer: e1edf9d1967ca96767dcc2b2d6df69f4`,
		469: `In a room N chairs are placed around a round table.
Knights enter the room one by one and choose at random an available empty
chair.
To have enough elbow room the knights always leave at least one empty
chair between each other.

When there aren't any suitable chairs left, the fraction C of empty chairs
is determined.
We also define E(N) as the expected value of C.
We can verify that E(4) = 1/2 and E(6) = 5/9.

Find E(10^18). Give your answer rounded to fourteen decimal places in the
form 0.abcdefghijklmn.


Answer: 3c2b641262880db5b735cfa4d4c957bc`,
	}

	for problem, want := range testProblems {
		if got := GetProblemText(problem); got != want {
			t.Errorf("GetProblemText(%d) = %s, want %s", problem, got, want)
		}
	}
}
