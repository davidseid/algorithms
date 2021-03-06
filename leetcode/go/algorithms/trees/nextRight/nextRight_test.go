package nextRight

import (
	"testing"

	"github.com/go-test/deep"
)

/*
116. Populating Next Right Pointers in Each Node

https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Follow up:

You may only use constant extra space.
Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.

Example 1:

Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
*/

/*
Rationale:
Use a Breadth First Search to link the nodes of each level. First initialize an array with the root,
then use a while loop as long as the array as nodes in it, iterate through the nodes and link them left to right,
then if the left or right are not nil, append them to the next queue. Repeate until there is nothing in the queue.

Time Complexity:
O(n), we must touch each node once in order to link it and access its children

Space Complexity:
O(n), we must build up a queue that holds the nodes at the current level, which results in n nodes

Followup:
Find a way to do this in constant space (possibly with recursion?)

Optimzations:
Implemented iterative solution using nested while loops, the first to track the left most node in the level, moving down one each time
The inner while loop builds the links and moves to the right across each level's children

This is now a constance space solution, still O(n) time complexity.

*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	leftMostInLevel := root

	for leftMostInLevel != nil && leftMostInLevel.Left != nil {
		curr := leftMostInLevel

		for curr != nil {
			curr.Left.Next = curr.Right

			if curr.Next == nil {
				curr.Right.Next = nil
			} else {
				curr.Right.Next = curr.Next.Left
			}

			curr = curr.Next
		}

		leftMostInLevel = leftMostInLevel.Left
	}

	return root
}

func TestConnect(t *testing.T) {
	t.Run("should connect next pointers", func(t *testing.T) {
		input := &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Left: &Node{
					Val: 6,
				},
				Right: &Node{
					Val: 7,
				},
			},
		}

		actual := connect(input)

		expected := &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Left: &Node{
					Val: 6,
				},
				Right: &Node{
					Val:  7,
					Next: nil,
				},
				Next: nil,
			},
			Next: nil,
		}

		expected.Left.Next = expected.Right
		expected.Left.Left.Next = expected.Left.Right
		expected.Left.Right.Next = expected.Right.Left
		expected.Right.Left.Next = expected.Right.Right

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
