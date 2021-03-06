package algorithms

import (
	"strconv"
	"testing"
)

/*
Source: https://leetcode.com/problems/decode-ways/

91. Decode Ways
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/

/*
Solved using recursing, time complexity is O(n^2), space complexity is O(n)

Follow up ideas:
- Can this be improved with go routines?
- What is the most efficient way to handle the bytes / strings / number conversions?
- How would we do this more statelessly?
*/

func TestNumDecodingsBasic(t *testing.T) {
	encoded := "12"

	expected := 2

	actual := numDecodings(encoded)

	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestNumDecodingsAdvanced(t *testing.T) {
	encoded := "226"

	expected := 3

	actual := numDecodings(encoded)

	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestNumDecodingsTricky(t *testing.T) {
	encoded := "01"

	expected := 0

	actual := numDecodings(encoded)

	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func numDecodings(s string) int {
	decodings := 0

	countDecodings(s, &decodings)

	return decodings
}

func countDecodings(s string, count *int) {
	if len(s) == 0 {
		*count++
		return
	}

	single := string(s[0])
	if single != "0" {
		countDecodings(string(s[1:]), count)
	} else {
		return
	}

	if len(s) > 1 {
		pair := string(s[0:2])

		intPair, err := strconv.Atoi(pair)

		if err != nil {
			panic("Invalid input")
		}

		if intPair >= 1 && intPair <= 26 {
			countDecodings(string(s[2:]), count)
		}
	}
}
