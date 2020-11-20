package validPalindrome

import (
	"strings"
	"testing"
)

/*
125. Valid Palindrome

https://leetcode.com/problems/valid-palindrome/

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:

Input: "race a car"
Output: false



Constraints:

s consists only of printable ASCII characters.
*/

/*
Time Complexity: O(n), we must touch every character in the worst case to know if the string is a palindrome
Space Complexity: O(1), we do not require extra space proportional to n
*/

func TestValidPalindrome(t *testing.T) {
	t.Run("should return true for palindromes", func(t *testing.T) {
		input := "A man, a plan, a canal: Panama"
		expected := true

		actual := isPalindrome(input)

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})

	t.Run("should return false for non palindromes", func(t *testing.T) {
		input := "race a car"
		expected := false

		actual := isPalindrome(input)

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})

	t.Run("should return true for palindromes with tricky characters", func(t *testing.T) {
		input := "Marge, let's \"[went].\" I await {news} telegram."
		expected := true

		actual := isPalindrome(input)

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(s) - 1

	for left <= right {
		leftChar := runes[left]
		rightChar := runes[right]

		if !isAlphanumeric(leftChar) {
			left++
			continue
		}

		if !isAlphanumeric(rightChar) {
			right--
			continue
		}

		if strings.ToLower(string(leftChar)) != strings.ToLower(string(rightChar)) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(char rune) bool {
	if (char >= 48 && char <= 57) || (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
		return true
	}

	return false
}
