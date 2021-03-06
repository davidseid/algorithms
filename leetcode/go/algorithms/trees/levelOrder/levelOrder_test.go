package levelorder

import (
	"testing"

	"github.com/go-test/deep"
)

/*
102. Binary Tree Level Order Traversal
https://leetcode.com/problems/binary-tree-level-order-traversal/
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestLevelOrder(t *testing.T) {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	expected := [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}

	actual := levelOrder(input)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}

	if root == nil {
		return levels
	}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		values := []int{}
		nextQueue := []*TreeNode{}

		for _, node := range queue {
			values = append(values, node.Val)

			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}

			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		queue = nextQueue
		levels = append(levels, values)
	}

	return levels
}
