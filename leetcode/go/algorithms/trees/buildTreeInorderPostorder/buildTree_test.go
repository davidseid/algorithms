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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	nextRoot := len(postorder) - 1
	return buildSubTree(0, len(inorder), &nextRoot, inorder, postorder)
}

func buildSubTree(start int, end int, nextRootIndex *int, inorder []int, postorder []int) *TreeNode {

	if start >= end || *nextRootIndex < 0 {
		return nil
	}

	root := &TreeNode{
		Val: postorder[*nextRootIndex],
	}

	var pivot int

	for i := start; i < end; i++ {
		if inorder[i] == root.Val {
			pivot = i
		}
	}

	*nextRootIndex--
	root.Right = buildSubTree(pivot+1, end, nextRootIndex, inorder, postorder)
	*nextRootIndex--
	root.Left = buildSubTree(start, pivot, nextRootIndex, inorder, postorder)

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
