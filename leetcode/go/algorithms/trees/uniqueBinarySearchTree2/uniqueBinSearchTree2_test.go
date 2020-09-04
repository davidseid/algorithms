package uniqueBinarySearchTree2

import (
	"testing"

	"github.com/go-test/deep"
)

/*
95. Unique Binary Search Trees II
https://leetcode.com/problems/unique-binary-search-trees-ii/

Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3


Constraints:

0 <= n <= 8
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	trees := []*TreeNode{}
	fakeRoot := &TreeNode{Val: -1}

	remaining := []int{}
	for i := 1; i <= n; i++ {
		remaining = append(remaining, i)
	}

	generateSubTrees(fakeRoot, fakeRoot, remaining, trees)
	return trees
}

func generateSubTrees(fakeRoot *TreeNode, node *TreeNode, remaining []int, trees []*TreeNode) {

}

func copyTree(original *TreeNode) *TreeNode {

}

func TestGenerateTrees(t *testing.T) {
	t.Run("should generate all unique BSTs", func(t *testing.T) {
		input := 3

		expected := []*TreeNode{
			&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
			&TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
			&TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
			&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		}

		actual := generateTrees(input)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
