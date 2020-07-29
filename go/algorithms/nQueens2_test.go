package algorithms

import "testing"

func TestNQueens2(t *testing.T) {

	input := 4
	expected := 2

	actual := solveNQueens(input)

	if actual != expected {
		t.Errorf("Got %d, wanted %d", actual, expected)
	}
}
