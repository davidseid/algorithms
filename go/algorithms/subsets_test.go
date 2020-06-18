package algorithms

import (
	"fmt"
	"reflect"
	"sort"
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

	result = getSubsets(nums, 0, currSubset, result)

	// sortResult(result)
	return result
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

func sortResult(matrix [][]int) {
	sort.Slice(matrix[:], func(i, j int) bool {
		for x := range matrix[i] {
			if matrix[i][x] == matrix[j][x] {
				continue
			}
			return matrix[i][x] < matrix[j][x]
		}
		return false
	})
}

func TestSubsets(t *testing.T) {
	input := []int{1, 2, 3}

	expected := [][]int{
		{3},
		{1},
		{2},
		{1, 2, 3},
		{1, 3},
		{2, 3},
		{1, 2},
		{},
	}

	actual := subsets(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

func TestSubsetsLonger(t *testing.T) {
	input := []int{9, 0, 3, 5, 7}

	expected := [][]int{
		{},
		{9},
		{0},
		{0, 9},
		{3},
		{3, 9},
		{0, 3},
		{0, 3, 9},
		{5},
		{5, 9},
		{0, 5},
		{0, 5, 9},
		{3, 5},
		{3, 5, 9},
		{0, 3, 5},
		{0, 3, 5, 9},
		{7},
		{7, 9},
		{0, 7},
		{0, 7, 9},
		{3, 7},
		{3, 7, 9},
		{0, 3, 7},
		{0, 3, 7, 9},
		{5, 7},
		{5, 7, 9},
		{0, 5, 7},
		{0, 5, 7, 9},
		{3, 5, 7},
		{3, 5, 7, 9},
		{0, 3, 5, 7},
		{0, 3, 5, 7, 9},
	}

	actual := subsets(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}
