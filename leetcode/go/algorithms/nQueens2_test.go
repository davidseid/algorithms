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

Rationale:
- Starting on first row, iterate through row
- For each column, try placing a queen, recurse, clean up
*/

func TestNQueens2(t *testing.T) {

	input := 4
	expected := 2

	actual := totalNQueens(input)

	if actual != expected {
		t.Errorf("Got %d, wanted %d", actual, expected)
	}
}

func totalNQueens(n int) int {
	board := makeBoard2(n)
	numSolutions := 0

	solveNQueens2(&numSolutions, &board, 0)
	return numSolutions
}

func solveNQueens2(numSolutions *int, pBoard *[][]int, row int) {
	board := *pBoard
	if row == len(board) {
		*numSolutions++
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid2(&board, row, col) {
			board[row][col] = 1
			solveNQueens2(numSolutions, &board, row+1)
			board[row][col] = 0
		}
	}
}

func isValid2(pBoard *[][]int, row int, col int) bool {
	board := *pBoard
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}

func makeBoard2(n int) [][]int {
	board := [][]int{}

	for i := 0; i < n; i++ {
		row := []int{}
		for i := 0; i < n; i++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}

	return board
}
