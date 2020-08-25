package subsets2

import (
	"testing"

	"github.com/go-test/deep"
)

/*
90. Subsets II
https://leetcode.com/problems/subsets-ii/

Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

func TestSubsetsWithDup(t *testing.T) {
	t.Run("should return the power set", func(t *testing.T) {
		input := []int{1, 2, 2}

		actual := subsetsWithDup(input)

		expected := [][]int{
			{2},
			{1},
			{1, 2, 2},
			{2, 2},
			{1, 2},
			{},
		}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func subsetsWithDup(nums []int) [][]int {

}
