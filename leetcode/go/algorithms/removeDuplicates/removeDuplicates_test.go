package removeduplicates

import (
	"testing"

	"github.com/go-test/deep"
)

/*
82. Remove Duplicates from Sorted List II
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Return the linked list sorted as well.

Example 1:

Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:

Input: 1->1->1->2->3
Output: 2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	return &ListNode{}
}

func buildLinkedListFromArray(arr []int) *ListNode {
	ll := &ListNode{}
	curr := ll

	for i := range arr {
		curr.Val = arr[i]
		if i < len(arr)-1 {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}

	return ll
}

func TestDeleteDuplicates(t *testing.T) {
	t.Run("should delete nodes with duplicate numbers from list", func(t *testing.T) {
		nums := []int{1, 2, 3, 3, 4, 4, 5}
		input := buildLinkedListFromArray(nums)

		expected := buildLinkedListFromArray([]int{1, 2, 5})

		actual := deleteDuplicates(input)
		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
