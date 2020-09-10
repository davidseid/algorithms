package binaryTreeLevelOrderTraversal

import (
	"testing"

	"github.com/go-test/deep"
)

/*
107. Binary Tree Level Order Traversal II

https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	levelOrder := [][]int{}

	if root == nil {
		return levelOrder
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		level := []int{}
		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append([][]int{level}, levelOrder...)
	}
	return levelOrder
}

func TestLevelOrderBottom(t *testing.T) {
	t.Run("should build bottom-up level order traversal node values", func(t *testing.T) {
		input := &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		}

		actual := levelOrderBottom(input)

		expected := [][]int{
			{15, 7},
			{9, 20},
			{3},
		}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
