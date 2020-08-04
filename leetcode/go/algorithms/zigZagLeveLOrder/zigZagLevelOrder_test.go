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
		[]int{20, 19},
		[]int{15, 7},
	}

	actual := zigZagLevelOrder(input)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
