package partitionList

import "testing"

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

}

func makeLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  arr[0],
		Next: &ListNode{},
	}

	curr := head.Next

	for _, v := range arr[1:] {
		curr.Val = v
		curr.Next = &ListNode{}
		curr = curr.Next
	}

	return head
}

func TestPartition(t *testing.T) {
	t.Run("should partition a linked list", func(t *testing.T) {

	})
}
