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

Rationale:
- Using a while loop to track left and right pointers, narrow in on the possible target.
- Because it is rotated, we must check to see if the left side is sorted, normal binary search,
if the left is unsorted, then that means the right is sorted, so we can check if it is on the right side,
and if so search right, otherwise search left
- Caveat is that because of duplicates, we may end up with the left, mid, and right being equal, in which case we can
narrow the range on each side by one until that is not the case.
*/

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := ((right - left) / 2) + left

		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
			continue
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
			continue
		}

		if nums[left] > nums[mid] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
			continue
		}
	}

	return false
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
