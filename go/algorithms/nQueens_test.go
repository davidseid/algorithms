package algorithms

import (
	"strconv"
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

	placeQueen(&board, 0, 0, n, &solutions)

	return solutions
}

func makeBoard(n int) []string {
	board := []string{}
	row := ""

	for i := 0; i < n; i++ {
		row += "0"
	}

	for i := 0; i < n; i++ {
		board = append(board, row)
	}

	return board
}

func placeQueen(board *[]string, row int, col int, queensRemaining int, solutions *[][]string) {
	if queensRemaining == 0 {
		cleanedBoard := cleanBoard(board)
		*solutions = append(*solutions, cleanedBoard)
		return
	}

	for r := row; r < len(*board); r++ {
		for c := col; c < len((*board)[0]); c++ {
			if spaceIsEmpty(board, r, c) {
				(*board)[r] = updateRow((*board)[r], c, "Q")
				markBoard(board, r, c)
				placeQueen(board, row+1, 0, queensRemaining-1, solutions)
				(*board)[r] = updateRow((*board)[r], c, "0")
				unmarkBoard(board, r, c)
			}
		}
	}

}

func updateRow(old string, col int, val string) string {
	newString := ""

	for i := 0; i < len(old); i++ {
		if i == col {
			newString += val
		} else {
			newString += string(old[i])
		}
	}

	return newString
}

func cleanBoard(board *[]string) []string {

	cleanedBoard := []string{}
	for row := range *board {
		cleanedRow := ""
		for col := 0; col < len((*board)[0]); col++ {
			if string((*board)[row][col]) == "Q" {
				cleanedRow += "Q"
			} else {
				cleanedRow += "."
			}
		}
		cleanedBoard = append(cleanedBoard, cleanedRow)
	}
	return cleanedBoard
}

func markBoard(board *[]string, row int, col int) {
	markHorizontal(board, row)
	markVertical(board, col)
	markDiagonals(board, row, col)
}

func markHorizontal(board *[]string, row int) {
	for col := 0; col < len((*board)[row]); col++ {
		if string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			current++
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
		}
	}
}

func markVertical(board *[]string, col int) {
	for row := range *board {
		if string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			current++
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
		}
	}
}

func markDiagonals(board *[]string, r int, c int) {
	col := c + 1
	for row := r + 1; row < len(*board); row++ {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {

			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			current++
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r + 1; row < len(*board); row++ {
		if col >= 0 && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			current++
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
			col--
		} else {
			break
		}
	}

	col = c + 1
	for row := r - 1; row >= 0; row-- {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			current++
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r - 1; row >= 0; row-- {
		if col >= 0 && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			current++
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
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
			current, err := strconv.Atoi(string((*board)[row][col]))
			if err != nil {
				panic(err)
			}
			if string((*board)[row][col]) != "0" {
				current--
			}
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
		}
	}
}

func unmarkVertical(board *[]string, col int) {
	for row := range *board {
		if string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			if string((*board)[row][col]) != "0" {
				current--
			}
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
		}
	}
}

func unmarkDiagonals(board *[]string, r int, c int) {
	col := c + 1
	for row := r + 1; row < len(*board); row++ {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			if string((*board)[row][col]) != "0" {
				current--
			}
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r + 1; row < len(*board); row++ {
		if col >= 0 && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			if string((*board)[row][col]) != "0" {
				current--
			}
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
			col--
		} else {
			break
		}
	}

	col = c + 1
	for row := r - 1; row >= 0; row-- {
		if col < len((*board)[0]) && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			if string((*board)[row][col]) != "0" {
				current--
			}
			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
			col++
		} else {
			break
		}
	}

	col = c - 1
	for row := r - 1; row >= 0; row-- {
		if col >= 0 && string((*board)[row][col]) != "Q" {
			current, err := strconv.Atoi(string((*board)[row][col]))

			if err != nil {
				panic(err)
			}
			if string((*board)[row][col]) != "0" {
				current--
			}

			(*board)[row] = updateRow((*board)[row], col, strconv.Itoa(current))
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
