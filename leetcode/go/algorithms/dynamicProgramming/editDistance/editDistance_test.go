package editDistance

import "testing"

/*
72. Edit Distance
https://leetcode.com/problems/edit-distance/

Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

func TestMinDistance(t *testing.T) {
	actual := minDistance("horse", "ros")

	if actual != 3 {
		t.Errorf("Got %d, want %d", actual, 3)
	}
}

func TestMinDistance2(t *testing.T) {
	actual := minDistance("intention", "execution")

	if actual != 5 {
		t.Errorf("Got %d, want %d", actual, 5)
	}
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
