package algorithms

import (
	"testing"
)

/*
101. Symmetric Tree
https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3
*/

/*
Approach 1: Recursive
Rationale:
- Recursively check if both sides are mirrors
- First check if the left and right of the root are mirrors
- To check if two nodes are mirrors, check if they are both the same,
and if their opposite children are mirrors.
- Time Complexity: O(n) with n = number of nodes
- Space Complexity: O(n) with n = number of nodes
*/

func TestIsSymmetric(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	actual := isSymmetric(tree)

	expected := true

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestIsSymmetric2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	actual := isSymmetric(tree)

	expected := false

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return areMirrors(root.Left, root.Right)
}

func areMirrors(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}

		if !areMirrors(left.Left, right.Right) {
			return false
		}

		if !areMirrors(left.Right, right.Left) {
			return false
		}

		return true
	}

	return false
}
