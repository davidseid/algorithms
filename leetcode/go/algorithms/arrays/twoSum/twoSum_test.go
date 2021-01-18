package twoSum

import (
	"testing"

	"github.com/go-test/deep"
)

/*
1. Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]



Constraints:

    2 <= nums.length <= 103
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    Only one valid answer exists.

*/

func twoSum(nums []int, target int) []int {
	counts := map[int][]int{}

	for i, num := range nums {
		if _, ok := counts[num]; ok {
			counts[num][1]++
			continue
		}

		counts[num] = []int{i, 1}
	}

	for i, num := range nums {
		counts[num][1]-- // to avoid reuse

		diff := target - num

		if tuple, ok := counts[diff]; ok && tuple[1] >= 1 {
			return []int{i, tuple[0]}
		}

		counts[num][1]++ // replace for next iteration
	}

	return []int{-1} // did not find twosum
}

func TestTwoSum(t *testing.T) {
	t.Run("should return indices of two numbers such that they add up to target", func(t *testing.T) {
		input := []int{2, 7, 11, 15}
		target := 9

		actual := twoSum(input, target)
		expected := []int{0, 1}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should return indices of two numbers such that they add up to target again", func(t *testing.T) {
		input := []int{3, 2, 4}
		target := 6

		actual := twoSum(input, target)
		expected := []int{1, 2}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should return indices of two numbers such that they add up to target with duplicates", func(t *testing.T) {
		input := []int{3, 3}
		target := 6

		actual := twoSum(input, target)
		expected := []int{0, 1}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
