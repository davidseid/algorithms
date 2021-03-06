package sameTree

import "testing"

/*
100. Same Tree
https://leetcode.com/problems/same-tree/

Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/

/*
Rationale:
Simple recursvie solution, if either of the nodes is nil, they both must be, otherwise return false
If neither are nil, if the values are the same and recursive calls to isSameTree on the left and right nodes of each
return true, then return true, otherwise return false

Time Complexity:
O(n), we iterate through the smallest of the trees fully.

Space Complexity:
O(n), due to recursive call stack, a recursive call for each node
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		}
		return false
	}

	if p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}

func TestSameTree(t *testing.T) {
	t.Run("should return true when two trees are identical", func(t *testing.T) {
		tree1 := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}

		tree2 := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}

		actual := isSameTree(tree1, tree2)

		expected := true

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})

	t.Run("should return false when trees are not the same", func(t *testing.T) {
		tree1 := &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		}

		tree2 := &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		}

		actual := isSameTree(tree1, tree2)

		expected := false

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})

	t.Run("should return false when trees are not the same 2", func(t *testing.T) {
		tree1 := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 1},
		}

		tree2 := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		}

		actual := isSameTree(tree1, tree2)

		expected := false

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})
}
