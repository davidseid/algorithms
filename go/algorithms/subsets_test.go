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

*/

func subsets(nums []int) [][]int {
	result := [][]int{}

	currSubset := []int{}

	getSubsets(currSubset, nums, result)
	return result
}

func getSubsets(currSubset []int, numsRemaining []int, subsets [][]int) [][]int {
	if len(numsRemaining) == 0 {
		return subsets
	}

	subsetWithNext := append(currSubset, numsRemaining[0])

	subsets = append(subsets, getSubsets(currSubset, numsRemaining[1:], subsets)...)
	subsets = append(subsets, getSubsets(subsetWithNext, numsRemaining[1:], subsets)...)
	return subsets
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
