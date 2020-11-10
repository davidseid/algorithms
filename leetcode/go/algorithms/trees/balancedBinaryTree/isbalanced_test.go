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

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	diff := leftHeight - rightHeight

	if math.Abs(float64(diff)) > 1 {
		return false
	}

	return true
}
