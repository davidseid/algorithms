package algorithms

import (
	"testing"
)

/*
98. Validate Binary Search Tree

Source: https://leetcode.com/problems/validate-binary-search-tree/

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Rationale:
- This is a recursive solution
- Time complexity is O(n), we must visit each node
- Space complexity is O(n) since we must keep up to the entire tree in the worst case

Followup:
- How about an iterative solution with a stack?
- How about using inorder traversal to simplify this?
- Can go routines help? Well they don't change the amount of resources needed obviously,
- but they do allow parallelization which could speed it up? Each call is a go routine
*/
func isValidBST(root *TreeNode) bool {
	return isValidSubtree(root, nil, nil)
}

func isValidSubtree(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}

	val := node.Val

	if min != nil && val <= *min {
		return false
	}

	if max != nil && val >= *max {
		return false
	}

	if !isValidSubtree(node.Right, &val, max) {
		return false
	}

	if !isValidSubtree(node.Left, min, &val) {
		return false
	}

	return true
}

func TestIsValidBSTBasic(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	expected := true

	actual := isValidBST(tree)

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestIsValidBSTAdvanced(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	expected := false

	actual := isValidBST(tree)

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestIsValidBSTMoreAdvanced(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	expected := true

	actual := isValidBST(tree)

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestIsValidBSTMoreMoreAdvanced(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	expected := false

	actual := isValidBST(tree)

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
