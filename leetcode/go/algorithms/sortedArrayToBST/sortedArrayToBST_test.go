package sortedArrayToBST

import (
	"testing"

	"github.com/go-test/deep"
)

/*
108. Convert Sorted Array to Binary Search Tree
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSortedArrayToBST(t *testing.T) {
	sortedArray := []int{-10, -3, 0, 5, 9}

	actual := sortedArrayToBST(sortedArray)

	expected := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
