package longestPalindromicSubstring

import "testing"

/*
5. Longest Palindromic Substring
Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"

Example 3:

Input: s = "a"
Output: "a"

Example 4:

Input: s = "ac"
Output: "a"

Constraints:
    1 <= s.length <= 1000
    s consist of only digits and English letters (lower-case and/or upper-case),
*/

func buildPalindrome(s string, left int, right int) string {
	for right < len(s) && left >= 0 && s[left] == s[right] {
		left--
		right++
	}

	return string(s[left+1 : right])
}

func updateLongestIfRequired(palindrome, longest string) string {
	if len(palindrome) > len(longest) {
		return palindrome
	}
	return longest
}

func longestPalindrome(s string) string {
	longest := ""

	for i := 0; i < len(s); i++ {
		palindrome := buildPalindrome(s, i, i)
		longest = updateLongestIfRequired(palindrome, longest)

		if i+1 < len(s) && s[i] == s[i+1] {
			palindrome := buildPalindrome(s, i, i+1)
			longest = updateLongestIfRequired(palindrome, longest)
		}
	}

	return longest
}

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Run("should check for longest palindromic substring in string", func(t *testing.T) {
		input := "babad"

		expected := "bab"

		actual := longestPalindrome(input)

		if actual != expected {
			t.Errorf("Got %s, want %s", actual, expected)
		}
	})

	t.Run("should check for longest palindromic substring in string", func(t *testing.T) {
		input := "cbbd"

		expected := "bb"

		actual := longestPalindrome(input)

		if actual != expected {
			t.Errorf("Got %s, want %s", actual, expected)
		}
	})

	t.Run("should check for longest palindromic substring in string with single character input", func(t *testing.T) {
		input := "a"

		expected := "a"

		actual := longestPalindrome(input)

		if actual != expected {
			t.Errorf("Got %s, want %s", actual, expected)
		}
	})

	t.Run("should check for longest palindromic substring in string when palindrome is single char", func(t *testing.T) {
		input := "ac"

		expected := "a"

		actual := longestPalindrome(input)

		if actual != expected {
			t.Errorf("Got %s, want %s", actual, expected)
		}
	})
}
