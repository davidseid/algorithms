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

Rationale:
In order to keep it balanced, recursively divide the sorted array in half along a pivot number.
The pivot number is the node value, and the left and right subtrees of the node use left
and right halves of the sorted array as inputs.

Time Complexity: O(n^2)
We build a node for each number in the sorted array input, however building a node involves making a new array for the left and right

Optimization: This can be improved if rather than copy the array, we just use start/end pointers on the initial array

Space Complexity: O(n^2)
We build left and right arrays for each node. This can also be optimized by using indices instead of copying the array
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

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, start int, end int) *TreeNode {
	if start >= end {
		return nil
	}

	mi := (end - start) / 2

	node := &TreeNode{
		Val: nums[mi],
	}

	node.Left = buildBST(nums, 0, mi)
	node.Right = buildBST(nums, mi+1, len(nums)-1)

	return node
}
