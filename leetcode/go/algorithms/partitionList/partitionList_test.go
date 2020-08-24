package partitionlist

import (
	"testing"

	"github.com/go-test/deep"
)

/*
86. Partition List
https://leetcode.com/problems/partition-list/

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	smallerHead := &ListNode{}
	biggerHead := &ListNode{}

	smaller := smallerHead
	bigger := biggerHead

	for head != nil {
		if head.Val < x {
			smaller.Next = head
			smaller = smaller.Next
		} else {
			bigger.Next = head
			bigger = bigger.Next
		}
		head = head.Next
	}
	smaller.Next = biggerHead.Next
	bigger.Next = nil
	return smallerHead.Next
}

func makeLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val: arr[0],
	}

	curr := head

	for _, v := range arr[1:] {
		curr.Next = &ListNode{
			Val: v,
		}
		curr = curr.Next
	}

	return head
}

func TestPartition(t *testing.T) {
	t.Run("should partition a linked list", func(t *testing.T) {
		input := makeLinkedList([]int{1, 4, 3, 2, 5, 2})

		actual := partition(input, 3)

		expected := makeLinkedList([]int{1, 2, 2, 4, 3, 5})

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
