package algorithms

import "testing"

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
	return true
}

func TestExists(t *testing.T) {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}

	word := "ABBCED"

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
	}

	word = "SEE"

	if !exist(board, word) {
		t.Errorf("Expected %s to exist", word)
	}

	word = "ABCB"

	if exist(board, word) {
		t.Errorf("Expected %s to not exist", word)
	}
}
