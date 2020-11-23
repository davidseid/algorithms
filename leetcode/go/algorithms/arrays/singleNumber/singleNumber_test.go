package singleNumber

import "testing"

/*
136. Single Number
https://leetcode.com/problems/single-number/

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?



Example 1:

Input: nums = [2,2,1]
Output: 1

Example 2:

Input: nums = [4,1,2,1,2]
Output: 4

Example 3:

Input: nums = [1]
Output: 1

Constraints:
    1 <= nums.length <= 3 * 104
    -3 * 104 <= nums[i] <= 3 * 104
    Each element in the array appears twice except for one element which appears only once.
*/

func TestSingleNumber(t *testing.T) {
	t.Run("should find single number when last", func(t *testing.T) {
		expected := 1
		actual := singleNumber([]int{2, 2, 1})

		if actual != expected {
			t.Errorf("Got %d, want %d", actual, expected)
		}
	})

	t.Run("should find single number when first", func(t *testing.T) {
		expected := 4
		actual := singleNumber([]int{4, 1, 2, 1, 2})

		if actual != expected {
			t.Errorf("Got %d, want %d", actual, expected)
		}
	})

	t.Run("should find single number when only", func(t *testing.T) {
		expected := 1
		actual := singleNumber([]int{1})

		if actual != expected {
			t.Errorf("Got %d, want %d", actual, expected)
		}
	})
}

func singleNumber(nums []int) int {

}
