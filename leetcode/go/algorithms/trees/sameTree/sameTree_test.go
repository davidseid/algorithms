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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

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
}
