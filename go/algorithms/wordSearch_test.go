package algorithms

import (
	"fmt"
	"testing"
)

/*

Source: https://leetcode.com/problems/word-search/

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells
are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.


Constraints:

board and word consists only of lowercase and uppercase English letters.
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/

func exist(board [][]byte, word string) bool {
	for row := range board {
		for col := range board[row] {
			if isPath(board, word, 0, row, col) == true {
				return true
			}
		}
	}

	return false
}

func isPath(board [][]byte, word string, i int, row int, col int) bool {
	if i >= len(word) {
		return true
	}

	if board[row][col] == word[i] {
		if i == len(word)-1 {
			return true
		}

		boardCopy := deepCopyMatrix(board)
		boardCopy[row][col] = 0 // blackout boardCopy spot to prevent repeats

		if row-1 >= 0 {
			if isPath(boardCopy, word, i+1, row-1, col) {
				return true
			}
		}

		if col+1 < len(boardCopy[row]) {
			if isPath(boardCopy, word, i+1, row, col+1) {
				return true
			}
		}

		if row+1 < len(boardCopy) {
			if isPath(boardCopy, word, i+1, row+1, col) {
				return true
			}
		}

		if col-1 >= 0 {
			if isPath(boardCopy, word, i+1, row, col-1) {
				return true
			}
		}
	}

	return false
}

func deepCopyMatrix(original [][]byte) [][]byte {
	matrixCopy := make([][]byte, len(original))

	for i := range original {
		matrixCopy[i] = make([]byte, len(original[i]))
		copy(matrixCopy[i], original[i])
	}

	return matrixCopy
}

func TestExists(t *testing.T) {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}

	word := "ABCCED"

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
	}

	board = [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}

	word = "SEE"

	fmt.Println(board)

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
	}

	board = [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}

	word = "ABCB"

	if exist(board, word) {
		t.Errorf("Expected %s to not exist", word)
	}

	board = [][]byte{
		[]byte("a"),
	}

	word = "a"

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
	}

	board = [][]byte{
		[]byte("CAA"),
		[]byte("AAA"),
		[]byte("BCD"),
	}

	word = "AAB"

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
	}
}
