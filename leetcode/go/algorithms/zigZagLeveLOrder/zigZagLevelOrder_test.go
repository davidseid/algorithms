package zigZagLeveLOrder

import (
	"testing"

	"github.com/go-test/deep"
)

/*
103. Binary Tree Zigzag Level Order Traversal
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestZigZagLevelOrder(t *testing.T) {
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
		[]int{20, 9},
		[]int{15, 7},
	}

	actual := zigzagLevelOrder(input)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestZigZagLevelOrder2(t *testing.T) {
	input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	expected := [][]int{
		[]int{1},
		[]int{3, 2},
		[]int{4, 5},
	}

	actual := zigzagLevelOrder(input)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	levels := [][]int{}

	if root == nil {
		return levels
	}

	queue := []*TreeNode{root}
	zig := true

	for len(queue) != 0 {
		levelOrder := []int{}
		nextQueue := []*TreeNode{}

		for _, node := range queue {
			levelOrder = append(levelOrder, node.Val)

			if zig == false {
				if node.Left != nil {
					nextQueue = append(nextQueue, node.Left)
				}

				if node.Right != nil {
					nextQueue = append(nextQueue, node.Right)
				}

			} else {
				if node.Right != nil {
					nextQueue = append(nextQueue, node.Right)
				}

				if node.Left != nil {
					nextQueue = append(nextQueue, node.Left)
				}
			}
		}

		levels = append(levels, levelOrder)
		queue = nextQueue
		zig = !zig
	}

	return levels
}
