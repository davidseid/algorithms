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

func solveNQueens(n int) [][]string {

	board := makeBoard(n)

	solutions := [][]string{}

	placeQueen(&board, 0, 0, n, &solutions)

	return solutions
}

func makeBoard(n int) []string {
	board := []string{}

	for i := 0; i < n; i++ {
		board = append(board, "0000")
	}

	return board
}

func placeQueen(board *[]string, row int, col int, queensRemaining int, solutions *[][]string) {
	if queensRemaining == 0 {
		cleanBoard(board)
		*solutions = append(*solutions, *board)
		return
	}

	for r := row; r < len(*board); r++ {
		for c := col; c < len((*board)[0]); c++ {
			if spaceIsEmpty(board, r, c) {
				(*board)[r] = (*board)[r][:c] + "Q" + (*board)[r][c+1:]
				markBoard(board, r, c)
				defer unmarkBoard(board, r, c)

				// call palceQueen on this
			}
		}
	}
}

func cleanBoard(board *[]string) {
	for row := range *board {
		for col := 0; col < len((*board)[0]); col++ {
			if string((*board)[row][col]) != "Q" {
				(*board)[row] = (*board)[row][:col] + "." + (*board)[row][col+1:]
			}
		}
	}
}

func markBoard(board *[]string, row int, col int) {
	markHorizontal(board, row)
	markVertical(board, col)
	markDiagonals(board, row, col)
}

func markHorizontal(board *[]string, row int) {
	for col := 0; col < len((*board)[row]); col++ {
		if string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current++
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
		}
	}
}

func markVertical(board *[]string, col int) {
	for row := range *board {
		if string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current++
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
		}
	}
}

func markDiagonals(board *[]string, r int, c int) {
	col := c + 1
	for row := r + 1; row < len(*board); row++ {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current++
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r + 1; row < len(*board); row++ {
		if col >= 0 && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current++
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col--
		} else {
			break
		}
	}

	col = c + 1
	for row := r - 1; row >= 0; row-- {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current++
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r - 1; row >= 0; row-- {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current++
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col--
		} else {
			break
		}
	}
}

func unmarkBoard(board *[]string, row int, col int) {
	unmarkHorizontal(board, row)
	unmarkVertical(board, col)
	unmarkDiagonals(board, row, col)
}

func unmarkHorizontal(board *[]string, row int) {
	for col := 0; col < len((*board)[row]); col++ {
		if string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current--
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
		}
	}
}

func unmarkVertical(board *[]string, col int) {
	for row := range *board {
		if string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current--
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
		}
	}
}

func unmarkDiagonals(board *[]string, r int, c int) {
	col := c + 1
	for row := r + 1; row < len(*board); row++ {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current--
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r + 1; row < len(*board); row++ {
		if col >= 0 && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current--
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col--
		} else {
			break
		}
	}

	col = c + 1
	for row := r - 1; row >= 0; row-- {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current--
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r - 1; row >= 0; row-- {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current := int((*board)[row][col])
			current--
			(*board)[row] = (*board)[row][:col] + string(current) + (*board)[row][col+1:]
			col--
		} else {
			break
		}
	}
}
func spaceIsEmpty(board *[]string, row int, col int) bool {
	if string((*board)[row][col]) == "0" {
		return true
	}

	return false
}

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
