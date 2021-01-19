package addtwonumbers

/*
2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumDummyHead := &ListNode{}
	pSum := sumDummyHead

	p1 := l1
	p2 := l2
	carry := 0

	for p1 != nil || p2 != nil {
		val1 := 0
		if p1 != nil {
			val1 = p1.Val
		}

		val2 := 0
		if p2 != nil {
			val2 = p2.Val
		}

		sum := val1 + val2 + carry
		carry = 0

		if sum > 9 {
			carry = 1
			sum = sum - 10
		}

		pSum.Next = &ListNode{
			Val: sum,
		}
		pSum = pSum.Next

		if p1 != nil {
			p1 = p1.Next
		}

		if p2 != nil {
			p2 = p2.Next
		}
	}

	if carry == 1 {
		pSum.Next = &ListNode{
			Val: 1,
		}
	}

	return sumDummyHead.Next
}
