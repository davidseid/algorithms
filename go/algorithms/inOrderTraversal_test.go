package algorithms

import (
	"reflect"
	"testing"
)

/**
94. Binary Tree Inorder Traversal
Source: https://leetcode.com/problems/binary-tree-inorder-traversal/

Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up:
- Recursive solution is trivial, could you do it iteratively?
- Can I do it more functionally by having traverse return arrays of ints
	and the main function accumulate them?

Solution Comments:
Trivial recursive solution.
Time complexity is O(n) with n being number of nodes because we must travel to every node
Space complexity is O(n) with n being the recursion depth, which in the worst case is equal to the number of nodes

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

func inorderTraversal(root *TreeNode) []int {
	traversed := []int{}

	if root != nil {
		traverse(root, &traversed)
	}

	return traversed
}

func traverse(node *TreeNode, traversed *[]int) {
	if node.Left != nil {
		traverse(node.Left, traversed)
	}

	*traversed = append(*traversed, node.Val)

	if node.Right != nil {
		traverse(node.Right, traversed)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestInOrderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	expected := []int{1, 3, 2}

	actual := inorderTraversal(tree)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestInOrderTraversalWithEmptyTree(t *testing.T) {
	var tree *TreeNode

	expected := []int{}

	actual := inorderTraversal(tree)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
