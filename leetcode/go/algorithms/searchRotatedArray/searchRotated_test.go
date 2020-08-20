package searchrotatedarray

import "testing"

/*
81. Search in Rotated Sorted Array II
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
Follow up:

This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?
*/

func search(nums []int, target int) bool {

}

func TestSearchRotatedSortedArray2(t *testing.T) {
	t.Run("should be true if target in array", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 0

		actual := search(nums, target)
		expected := true

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})

	t.Run("should be false if target not in array", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 3

		actual := search(nums, target)
		expected := false

		if actual != expected {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})
}
