package algorithms

import (
	"fmt"
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

/*
Sliding Window Algorithm:
- Maintain two pointers, right that expands the window and left that contracts
- Expand right pointer until the window contains all the necessary characters, then contract until it doesn't
- Update the minimum window as we go
- Requires a map containing all the required chars
- Requires a map of the count of required chars in current window
- Requires a minWindow variable which is eventually returned
*/

func minWindow(s string, t string) string {
	requiredCharacterCounts := map[string]int{}

	for _, char := range t {
		requiredChar := string(char)
		if count, ok := requiredCharacterCounts[requiredChar]; !ok {
			fmt.Println(count)
			requiredCharacterCounts[requiredChar] = 1
		} else {
			requiredCharacterCounts[requiredChar] = count + 1
		}
	}

	isValidWindow := false
	characterCountsInWindow := map[string]int{}
	minWindow := ""
	left := 0
	right := 0

	for right < len(s) {
		rightChar := string(s[right])
		if _, ok := requiredCharacterCounts[rightChar]; ok {
			characterCountsInWindow[rightChar] = characterCountsInWindow[rightChar] + 1
			isNowValid := true

			for requiredChar, requiredCount := range requiredCharacterCounts {
				if characterCountsInWindow[requiredChar] < requiredCount {
					isNowValid = false
				}
			}

			isValidWindow = isNowValid

			for isValidWindow && left <= right {
				var window string

				if right+1 == len(s) {
					window = s[left:]
				} else {
					window = s[left : right+1]
				}

				if len(minWindow) == 0 || len(window) < len(minWindow) {
					minWindow = window
				}
				leftChar := string(s[left])
				if _, ok := requiredCharacterCounts[leftChar]; ok {
					characterCountsInWindow[leftChar] = characterCountsInWindow[leftChar] - 1
					if characterCountsInWindow[leftChar] < requiredCharacterCounts[leftChar] {
						isValidWindow = false
					}
				}
				left++
			}
		}
		right++
	}

	return minWindow
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

func TestMinWindowWithDuplicateRequiredChars(test *testing.T) {
	s := "a"
	t := "aa"

	expected := ""

	actual := minWindow(s, t)

	if actual != expected {
		test.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestMinWindowWithStringEqualToRequiredChars(test *testing.T) {
	s := "a"
	t := "a"

	expected := "a"

	actual := minWindow(s, t)

	if actual != expected {
		test.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestMinWindowWithStringEqualToRequiredCharsWithDuplicates(test *testing.T) {
	s := "aa"
	t := "aa"

	expected := "aa"

	actual := minWindow(s, t)

	if actual != expected {
		test.Errorf("Got %s, want %s", actual, expected)
	}
}
