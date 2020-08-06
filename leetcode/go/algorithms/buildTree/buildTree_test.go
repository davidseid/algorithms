package buildTree

import (
	"testing"

	"github.com/go-test/deep"
)

/*
105. Construct Binary Tree from Preorder and Inorder Traversal
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7

Rationale:
- Start with first in preorder as root
- Use current preorder value to find all elements
  to left and right in inorder array.
- Recursively build subtrees using the next inorder value
- If the left preorder array is not empty, use the next inorder on the left
- If the right preorder array is not empty, use the next inorder on the right
- Base case is when the left and right arrays are empty, meaning that node has no children
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	expected := &TreeNode{
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

	actual := buildTree(preorder, inorder)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	initialPreorderIndex := 0

	root := &TreeNode{
		Val: preorder[initialPreorderIndex],
	}

	buildSubtrees(root, &initialPreorderIndex, preorder, inorder)

	return root
}

func buildSubtrees(node *TreeNode, poIndex *int, preorder []int, inorder []int) {
	left, right := getLeftAndRight(node.Val, inorder)

	if len(left) != 0 {
		*poIndex = *poIndex + 1
		node.Left = &TreeNode{
			Val: preorder[*poIndex],
		}
		buildSubtrees(node.Left, poIndex, preorder, left)
	}

	if len(right) != 0 {
		*poIndex = *poIndex + 1
		node.Right = &TreeNode{
			Val: preorder[*poIndex],
		}
		buildSubtrees(node.Right, poIndex, preorder, right)
	}
}

func getLeftAndRight(nodeValue int, inorder []int) ([]int, []int) {
	var targetIndex int

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == nodeValue {
			targetIndex = i
			break
		}
	}

	right := []int{}

	if targetIndex < len(inorder)-1 {
		right = inorder[targetIndex+1:]
	}

	return inorder[:targetIndex], right
}
