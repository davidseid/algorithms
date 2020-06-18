package algorithms

import (
	"fmt"
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

*/

func subsets(nums []int) [][]int {
	result := [][]int{}
	currSubset := []int{}

	return getSubsets(nums, 0, currSubset, result)
}

func getSubsets(nums []int, index int, curr []int, subsets [][]int) [][]int {
	if index == len(nums) {
		fmt.Println(curr)
		subsets = append(subsets, curr)
		return subsets
	}

	currWithNext := make([]int, len(curr))
	copy(currWithNext, curr)
	currWithNext = append(currWithNext, nums[index])

	subsets = getSubsets(nums, index+1, currWithNext, subsets)
	subsets = getSubsets(nums, index+1, curr, subsets)
	return subsets
}

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
