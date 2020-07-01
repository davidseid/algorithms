package algorithms

import (
	"testing"
)

/*
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) string {
	charToIndex := map[string]int{}

	for _, v := range t {
		charToIndex[string(v)] = -1
	}

	windows := [][]int{}
	indexToChar := map[int]string{}
	charsFound := 0

	i := 0

	for i < len(s) {
		currChar := string(s[i])
		if charIndex, ok := charToIndex[currChar]; ok {

			if charIndex == -1 {
				// add it to the map
				// add it to the indexToChar
				// increment end
			} else {
				// get old index
				// update charToIndex
				// update indexToChar (old and new)
			}

			// if charsFound == len(t)
			// add window based on indexToChar
		}
		i++
	}

	return ""
}

func TestMinWindow(test *testing.T) {
	s := "ADOBECODEBANC"
	t := "ABC"

	expected := "BANC"

	actual := minWindow(s, t)

	if actual != expected {
		test.Errorf("Got %s, want %s", actual, expected)
	}
}
