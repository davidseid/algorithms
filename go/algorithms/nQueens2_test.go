package algorithms

import "testing"

/*
Problem:
https://leetcode.com/problems/n-queens-ii/
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

Example:

Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
*/

func TestNQueens2(t *testing.T) {

	input := 4
	expected := 2

	actual := solveNQueens(input)

	if actual != expected {
		t.Errorf("Got %d, wanted %d", actual, expected)
	}
}

func totalNQueens(n int) int {

}
