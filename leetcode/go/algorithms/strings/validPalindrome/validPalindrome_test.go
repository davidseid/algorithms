package validPalindrome

import (
	"fmt"
	"regexp"
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
}

func isPalindrome(s string) bool {
	regex, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		panic("Failed to compile regex")
	}

	s = strings.ToLower(s)
	s = regex.ReplaceAllString(s, "")

	left := 0
	right := len(s) - 1

	for left <= right {
		fmt.Println(string(s[left]))
		fmt.Println(string(s[right]))

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
