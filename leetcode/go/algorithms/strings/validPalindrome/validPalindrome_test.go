package validPalindrome

import "testing"

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
}
