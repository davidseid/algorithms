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

/*
Rationale:
Breadth first search to traverse nodes from top to bottom by level, assembling the nodes for each level using an iterative approach
with a custom queue data structure.

Then simply reverse the levels to get a bottom up list
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	nodes []*TreeNode
}

func (q *Queue) Enqueue(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Dequeue() *TreeNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.nodes)
}

func (q *Queue) enqueueAnyChildren(node *TreeNode) {
	if node.Left != nil {
		q.Enqueue(node.Left)
	}
	if node.Right != nil {
		q.Enqueue(node.Right)
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	levelOrder := [][]int{}

	q := initializeQueue(root)

	for q.Size() > 0 {
		size := q.Size()
		level := []int{}
		for i := 0; i < size; i++ {
			node := q.Dequeue()
			level = append(level, node.Val)
			q.enqueueAnyChildren(node)
		}
		levelOrder = append(levelOrder, level)
	}

	reverseLevels(levelOrder)
	return levelOrder
}

func initializeQueue(root *TreeNode) Queue {
	q := Queue{}

	if root != nil {
		q.Enqueue(root)
	}
	return q
}

func reverseLevels(levelOrder [][]int) {
	for i, j := 0, len(levelOrder)-1; i < j; i, j = i+1, j-1 {
		levelOrder[i], levelOrder[j] = levelOrder[j], levelOrder[i]
	}
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
