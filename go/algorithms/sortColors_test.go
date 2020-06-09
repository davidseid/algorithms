package algorithms

import (
	"reflect"
	"testing"
)

/*
Source: https://leetcode.com/problems/sort-colors/

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
*/

func sortColors(nums []int) {
	low := 0
	mid := len(nums) / 2
	high := len(nums)

	curr := 0

	for currentIndexIsLessThanHighIndex(curr, high) {
		if currentValueIsLessThanMidValue(curr, mid, nums) {
			swap(low, curr, nums)
			low++
			curr++
			continue
		}

		if currentValueIsGreaterThanMidValue(curr, mid, nums) {
			high--
			swap(high, curr, nums)
			continue
		}

		curr++
	}
}

func currentIndexIsLessThanHighIndex(curr int, high int) bool {
	return curr < high
}

func currentValueIsLessThanMidValue(curr int, mid int, nums []int) bool {
	return nums[curr] < nums[mid]
}

func currentValueIsGreaterThanMidValue(curr int, mid int, nums []int) bool {
	return nums[curr] > nums[mid]
}

func swap(a int, b int, nums []int) {
	nums[a], nums[b] = nums[b], nums[a]
}

// Naive solution, two pass and constant space

// type counts struct {
// 	Reds   int
// 	Whites int
// 	Blues  int
// }

// const (
// 	red = iota
// 	white
// 	blue
// )
// func sortColors(nums []int) {

// 	counts := countColors(nums)

// 	orderColors(nums, counts)
// }

// func countColors(nums []int) counts {
// 	var counts = counts{}

// 	for _, color := range nums {
// 		if color == red {
// 			counts.Reds++
// 		}

// 		if color == white {
// 			counts.Whites++
// 		}

// 		if color == blue {
// 			counts.Blues++
// 		}
// 	}

// 	return counts
// }

// func orderColors(nums []int, counts counts) {
// 	for i := range nums {
// 		if counts.Reds > 0 {
// 			nums[i] = red
// 			counts.Reds--
// 			continue
// 		}

// 		if counts.Whites > 0 {
// 			nums[i] = white
// 			counts.Whites--
// 			continue
// 		}

// 		nums[i] = blue
// 		counts.Blues--
// 	}
// }

func TestSortColors(t *testing.T) {
	actual := []int{2, 0, 2, 1, 1, 0}
	expected := []int{0, 0, 1, 1, 2, 2}

	sortColors(actual)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}
