package algorithms

import (
	"testing"
)

/*
101. Symmetric Tree
https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3
*/

/*
Approach 1: Recursive
Rationale:
- Recursively check if both sides are mirrors
- First check if the left and right of the root are mirrors
- To check if two nodes are mirrors, check if they are both the same,
and if their opposite children are mirrors.
- Time Complexity: O(n) with n = number of nodes
- Space Complexity: O(n) with n = number of nodes
*/

/*
Approach 2: Iterative
Rationale:
- Iteratively check for symmetry using a queue to maintain order,
similar to BFS of a tree, except that we add the nodes in odd orders
- This implementation uses an inefficent queue that uses a slice
- This is inefficient because slices/arrays require all indexes to shift
when appending to the front of the queue. A linked list is more optimal.

*/

func TestIsSymmetric(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	actual := isSymmetric(tree)

	expected := true

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestIsSymmetric2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	actual := isSymmetric(tree)

	expected := false

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

type Queue struct {
	Values []*TreeNode
}

func (q *Queue) Enqueue(val *TreeNode) {
	q.Values = append([]*TreeNode{val}, q.Values...)
}

func (q *Queue) Dequeue() *TreeNode {
	dequeued := q.Values[len(q.Values)-1]
	q.Values = q.Values[:len(q.Values)-1]
	return dequeued
}

func (q *Queue) IsEmpty() bool {
	if len(q.Values) == 0 {
		return true
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	queue := Queue{}

	queue.Enqueue(root)
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		left := queue.Dequeue()
		right := queue.Dequeue()

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue.Enqueue(left.Left)
		queue.Enqueue(right.Right)
		queue.Enqueue(left.Right)
		queue.Enqueue(right.Left)
	}

	return true
}

// Recursive
// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	return areMirrors(root.Left, root.Right)
// }

// func areMirrors(left *TreeNode, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	}

// 	if left != nil && right != nil {
// 		if left.Val != right.Val {
// 			return false
// 		}

// 		if !areMirrors(left.Left, right.Right) {
// 			return false
// 		}

// 		if !areMirrors(left.Right, right.Left) {
// 			return false
// 		}

// 		return true
// 	}

// 	return false
// }
