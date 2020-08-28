package reverseLinkedList2

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

/*
92. Reverse Linked List II
https://leetcode.com/problems/reverse-linked-list-ii/

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if m == n {
		return head
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head

	var beforeMNode *ListNode
	var mNode *ListNode

	curr := dummyHead
	index := 0

	for curr != nil && index < m-1 {
		curr = curr.Next
		index++
	}

	beforeMNode = curr
	mNode = curr.Next

	prev := mNode
	curr = mNode.Next
	index = m + 1

	for curr != nil && index <= n {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next

		index++
	}

	mNode.Next = curr
	beforeMNode.Next = prev

	return dummyHead.Next
}

func buildLinkedList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead

	for _, v := range nums {
		curr.Next = &ListNode{
			Val: v,
		}
		curr = curr.Next
	}

	return dummyHead.Next
}

func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr)
		curr = curr.Next
	}
}

func TestReverseBetween(t *testing.T) {
	input := buildLinkedList([]int{1, 2, 3, 4, 5})

	actual := reverseBetween(input, 2, 4)
	printLinkedList(actual)

	expected := buildLinkedList([]int{1, 4, 3, 2, 5})

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
