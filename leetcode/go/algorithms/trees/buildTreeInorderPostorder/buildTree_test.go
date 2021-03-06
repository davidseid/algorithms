package buildTreeInorderPostorder

import (
	"testing"

	"github.com/go-test/deep"
)

/*
106. Construct Binary Tree from Inorder and Postorder Traversal

https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

/*
Rationale:
The last value in the post order is the root.

Recursively, grab the next root from the post order, use that to divide up the
inorder and recursively call again. The base case is when the start > end or the nextRoot is greater than the
length of the post order.

Time Complexity: O(n) with n being the length of the post/inorder arrays
Space Complexity: O(n) with n being the length of the post/inorder arrays, since we need to build the tree

Followup:
Need to redo this problem since I failed to catch the key insight of identifying the root from the post order list,
also had some trouble with the implementation, particularly around keeping the index in range.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildSubTree(0, len(inorder)-1, len(postorder)-1, inorder, postorder)
}

func buildSubTree(start int, end int, nextRootIndex int, inorder []int, postorder []int) *TreeNode {
	if nextRootIndex >= len(postorder) || start > end {
		return nil
	}

	root := &TreeNode{
		Val: postorder[nextRootIndex],
	}

	pivot := 0

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			pivot = i
		}
	}

	root.Left = buildSubTree(start, pivot-1, nextRootIndex-1-end+pivot, inorder, postorder)
	root.Right = buildSubTree(pivot+1, end, nextRootIndex-1, inorder, postorder)

	return root
}

func TestBuildTree(t *testing.T) {
	t.Run("should build a binary tree from inorder and postorder arrays", func(t *testing.T) {
		inorder := []int{9, 3, 15, 20, 7}
		postorder := []int{9, 15, 7, 20, 3}

		actual := buildTree(inorder, postorder)

		expected := &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
