package uniqueBinarySearchTree

import "testing"

/*
96. Unique Binary Search Trees
https://leetcode.com/problems/unique-binary-search-trees/

Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:

Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3


Constraints:

1 <= n <= 19
*/

/*
Rationale:
Solved with recursive catalan formula. Had to look it up though, need to resolve with
more intuitive solution that does not involve building all trees manually and deduping them.
*/

func TestUniqueBinarySearchTree(t *testing.T) {
	input := 3

	actual := numTrees(input)

	expected := 5

	if actual != expected {
		t.Errorf("Got %d, want %d", actual, expected)
	}
}

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	count := 0
	for i := 0; i < n; i++ {
		count += numTrees(i) * numTrees(n-1-i)
	}

	return count
}
