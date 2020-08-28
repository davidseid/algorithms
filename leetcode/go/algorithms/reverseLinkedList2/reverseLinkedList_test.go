package reverseLinkedList2

import (
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

func TestReverseBetween(t *testing.T) {
	input := buildLinkedList([]int{1, 2, 3, 4, 5})

	actual := reverseBetween(input, 2, 4)

	expected := buildLinkedList([]int{1, 4, 3, 2, 5})

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
