package algorithms

import (
	"fmt"
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValidSubtree(root, nil, nil)
}

func isValidSubtree(node *TreeNode, min *int, max *int) bool {
	fmt.Println("Current Node", *node)
	if min != nil {
		fmt.Println("min", *min)
	}

	if max != nil {
		fmt.Println("max", *max)
	}

	if node.Left != nil {
		if node.Val <= node.Left.Val {
			return false
		}

		if min != nil && node.Left.Val <= *min {
			return false
		}

		if max == nil {
			max = &node.Val
		}

		if !isValidSubtree(node.Left, nil, max) {
			return false
		}
	}

	fmt.Println("Left side is valid for node", *node)

	if node.Right != nil {
		if node.Val >= node.Right.Val {
			return false
		}

		if max != nil && node.Right.Val >= *max {
			return false
		}

		if min == nil {
			min = &node.Val
		}

		if !isValidSubtree(node.Right, min, nil) {
			return false
		}
	}

	fmt.Println("Right side is valid for node", *node)

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
