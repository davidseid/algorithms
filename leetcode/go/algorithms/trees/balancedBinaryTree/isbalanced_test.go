package balancedBinaryTree

import (
	"math"
	"testing"
)

/*
110. Balanced Binary Tree
https://leetcode.com/problems/balanced-binary-tree/

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: true
Example 2:


Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true


Constraints:

The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104
*/

/*
Time Complexity: O(nlogn)
If n is the number of nodes, there are logn levels in the tree. For each node, we traverse logn nodes to calculate the depth.

Space Complexity: O(1)
We do not require any extra space with this approach, we are simply using CPU to calculate the depths.

Improvement ideas:
1) Add memoization. In the current solution, we recount the depth of nodes many times, wasting CPU and time. If we append a depth value to the
nodes or otherwise maintain the previously calculated depths, we can significantly reduce the time complexity, with a mild tradeoff for space.

2) Try BFS. Currently, because we are evaluating in a depth first manner. This means we can have a very long left subtree from the root and no right subtree,
but from the algorithm's point of view we must first exhaust the depth calculations on the left subtree. This may not necessarily be an improvement,
it is unclear if it is just trading off which we calculate first. Is that any more efficient than DFS? Imagine other trees that are more performant for DFS...
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestIsBalanced(t *testing.T) {
	t.Run("should return true if left and right subtrees of every node differ in height by no more than one", func(t *testing.T) {
		root := &TreeNode{
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

		if isBalanced(root) != true {
			t.Errorf("got false, expected true")
		}
	})

	t.Run("should return false if left and right subtrees of every node differ in height by more than one", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		}

		if isBalanced(root) != false {
			t.Errorf("got true, expected false")
		}
	})
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)

	diff := math.Abs(float64(leftHeight - rightHeight))

	if diff > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := depth(root.Left) + 1
	rightDepth := depth(root.Right) + 1

	return int(math.Max(float64(leftDepth), float64(rightDepth)))
}
