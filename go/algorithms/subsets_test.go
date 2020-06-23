package algorithms

import (
	"reflect"
	"testing"
)

/*
Source: https://leetcode.com/problems/subsets/

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.
*/

/*
Rationale:
For each element in the input, fork a subset that includes it and one that does not
This can be done recursively with the base case being that there are no more elements in the input to iterate over

Example:
Begin with 1, 2, 3 and an empty subset

Call the power set again on 2, 3 with a subset of {} and a subset of {1}
The next time the first call will result in {}, {2}, the second will be {1},{1,2} and so on

Lessons learned:
- Go's built-in append function manipulates the underlying array even when the result assigned to a new variable.
- Better to make an empty array of the appropriate length, use the built-in copy(dest,src) to fill it, and then use append
- Sorting nested arrays can be done by implementing the Sort interface for a matrix

Time complexity: O(N*2^n) Exponential because we have two branches for each element and we have to make a copy each time.
Space complexity: O(N*2^n) Exponential because the total storage is represented by the size of the output and we are copying the result each time.
However, in terms of auxiliary space it could be possible to optimize this to use a bitmap / integer to hold the values...

Followup:
- Rewrite recursive solution
- Rewrite bitmap solution
*/

// REWRITE RECURSIVE SOLUTION
func subsets(nums []int) [][]int {
	return [][]int{[]int{}}
}

// // ITERATIVE SOLUTION
// func subsets(nums []int) [][]int {
// 	results := [][]int{
// 		[]int{},
// 	}

// 	for _, num := range nums {
// 		nextSubsets := deepCopy(results)

// 		for _, subset := range nextSubsets {
// 			subset = append(subset, num)
// 			results = append(results, subset)
// 		}
// 	}

// 	return results
// }

// func deepCopy(src [][]int) [][]int {
// 	nestedCopy := make([][]int, len(src))

// 	for i, v := range src {
// 		dupe := make([]int, len(v))
// 		copy(dupe, v)
// 		nestedCopy[i] = dupe
// 	}

// 	return nestedCopy
// }

// BITMASK SOLUTION
// func subsets(nums []int) [][]int {
// 	n := float64(len(nums))
// 	output := [][]int{}

// 	for i := math.Pow(2, n); i < math.Pow(2, n+1); i++ {
// 		bitmask := bin(i)[1:]
// 		subset := []int{}

// 		for j := 0; j < int(n); j++ {
// 			if len(bitmask) > j && bitmask[j] == '1' {
// 				subset = append(subset, nums[j])
// 			}
// 		}
// 		output = append(output, subset)
// 	}

// 	return output
// }

// func bin(num float64) string {
// 	n := int64(num)
// 	return strconv.FormatInt(n, 2)
// }

// RECURSIVE SOLUTION
// func subsets(nums []int) [][]int {
// 	result := [][]int{}
// 	currSubset := []int{}

// 	return getSubsets(nums, 0, currSubset, result)
// }

// func getSubsets(nums []int, index int, curr []int, subsets [][]int) [][]int {
// 	if index == len(nums) {
// 		subsets = append(subsets, curr)
// 		return subsets
// 	}

// 	currWithNext := make([]int, len(curr))
// 	copy(currWithNext, curr)
// 	currWithNext = append(currWithNext, nums[index])

// 	subsets = getSubsets(nums, index+1, currWithNext, subsets)
// 	subsets = getSubsets(nums, index+1, curr, subsets)
// 	return subsets
// }

func TestSubsets(t *testing.T) {
	input := []int{1, 2, 3}

	expected := [][]int{
		{1, 2, 3},
		{1, 2},
		{1, 3},
		{1},
		{2, 3},
		{2},
		{3},
		{},
	}

	actual := subsets(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}
