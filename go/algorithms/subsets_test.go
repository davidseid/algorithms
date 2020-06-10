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
Initialize result with empty value.
For each number in input
	loop again, add an entry each time
*/

func subsets(nums []int) [][]int {
	result := [][]int{
		{},
	}

	return result
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
