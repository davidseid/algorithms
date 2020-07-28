package algorithms

import (
	"testing"

	"github.com/go-test/deep"
)

/*
51. N-Queens
https://leetcode.com/problems/n-queens/
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
*/

/*
Rationale:
- This is an initial implementation using recursion and backtracking.
A simpler to implement algorithm would be to copy the board on each recursive step,
but that would be less space efficient. However, the board marking and unmarking
is very inefficient in this implementation because it uses an array of strings as the board,
indexing into strings and modifying strings is inefficient because they are immutable in golang.

This was done because the answer is expected as an array of strings, however it is not worth the
inefficiency, considering I'm using numbers to track the marking anyway.

Updated solution rationale:
- Start on the first row, for each column in that row, try to place a queen
- if can place a queen, recurse in with that board and the next column.
- when you come out, backtrack
- the base case is when you've gone past the last row
- A key insight here is that if you must place n queens on an n*n board, there must be exactly
one queen per row.
- Because we only go row by row, we only need to check three things to know if it is a valid spot:
-- is the column in the rows leading up to this open?
-- is the 45 degree angle up from the spot open?
-- is the 135 degree angle up/left from the spot open?
- Abandons the marking system, we just place the queen, and check for validity when we need to.

Some followup optimizations to explore:
- Use [][]int data structure instead, and only convert to a solution format at the end.
- Try to use a different system of marking, do we need to keep track of the numbers in this way
or can we just scan if it is safe before placing.
- Do we even need to scan the whole board or can we do only part of the board
- How can we take advantage of mirroring/rotating the solution to find additional solutions?

Goal is to speed it up enough to pass leetcode test
*/

func TestSolveNQueens(t *testing.T) {
	input := 4

	expected := [][]string{
		[]string{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		},
		[]string{
			"..Q.",
			"Q...",
			"...Q",
			".Q..",
		},
	}

	actual := solveNQueens(input)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func solveNQueens(n int) [][]string {
	board := makeBoard(n)

	solutions := [][]string{}

	solve(&solutions, &board, 0)
	return solutions
}

func solve(solutions *[][]string, pBoard *[][]int, row int) {
	board := *pBoard
	if row == len(board) {
		constructed := construct(board)
		*solutions = append(*solutions, constructed)
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(&board, row, col) {
			board[row][col] = 1
			solve(solutions, &board, row+1)
			board[row][col] = 0
		}
	}
}

func isValid(pBoard *[][]int, row int, col int) bool {
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

// convert to simple copy
func construct(board [][]int) []string {
	path := []string{}
	for i := range board {
		row := ""
		for j := range board[i] {
			if board[i][j] == 1 {
				row += "Q"
			} else {
				row += "."
			}
		}
		path = append(path, row)
	}

	return path
}

func makeBoard(n int) [][]int {
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
