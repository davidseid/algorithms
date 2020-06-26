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

Followup:
Optimize by backup cleanup rather than copying matrix each recurse (CHECK)
Optimize by changing base case. Rather than attempt to check path on each side,
if an adjacent square matches, if it ends the word just return, otherwise, continue
*/

func exist(board [][]byte, word string) bool {
	for row := range board {
		for col := range board[row] {
			if board[row][col] == word[0] {
				if len(word) == 1 || isPath(board, word, 0, row, col) == true {
					return true
				}
			}
		}
	}

	return false
}

func isPath(board [][]byte, word string, i int, row int, col int) bool {
	originalValue := board[row][col]
	board[row][col] = 0
	i++

	defer func() {
		board[row][col] = originalValue
	}()

	if row-1 >= 0 {
		if i == len(word)-1 && board[row-1][col] == word[i] {
			return true
		}

		if i < len(word)-1 && board[row-1][col] == word[i] {
			if isPath(board, word, i, row-1, col) {
				return true
			}
		}
	}

	if row+1 < len(board) {
		if i == len(word)-1 && board[row+1][col] == word[i] {
			return true
		}

		if i < len(word)-1 && board[row+1][col] == word[i] {
			if isPath(board, word, i, row+1, col) {
				return true
			}
		}
	}

	if col-1 >= 0 {
		if i == len(word)-1 && board[row][col-1] == word[i] {
			return true
		}

		if i < len(word)-1 && board[row][col-1] == word[i] {
			if isPath(board, word, i, row, col-1) {
				return true
			}
		}
	}

	if col+1 < len(board[row]) {
		if i == len(word)-1 && board[row][col+1] == word[i] {
			return true
		}

		if i < len(word)-1 && board[row][col+1] == word[i] {
			if isPath(board, word, i, row, col+1) {
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

	word = "SEE"

	fmt.Println(board)

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
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
