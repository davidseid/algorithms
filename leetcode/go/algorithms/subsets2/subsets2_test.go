package subsets2

import (
	"fmt"
	"sort"
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

Rationale:
Recursively collect subsets by either choosing or not choosing each value.
Hacky way to avoid dupes is to maintain a map of "seen" values, unfortunately requiring
we stringify each subset into a key for lookup prior to storing. This is likely not an efficient approach.
We also must sort the array prior to collecting the subsets, otherwise we would have to sort each subset prior
to checking the key.

Time Complexity:
O(n * 2^n), we have two branches with a depth of n. The O(nlogn) initial sort is neglible, however we also iterate
through each subset to create a key, which add significant complexity.

Space Complexity:
O(n * 2^n), same reason as above, the recursive calls and the n-sized map.

Followup:
There is probably a clever way to ensure we don't run into dupes without having to use the map. We may not be able to avoid
the initial sort, but perhaps that combined with an algorithm change regarding when we recurse could lead to a O(2^n) solution.
*/

func TestSubsetsWithDup(t *testing.T) {
	t.Run("should return the power set", func(t *testing.T) {
		input := []int{1, 2, 2}

		actual := subsetsWithDup(input)

		fmt.Println(actual)

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

	subsets := [][]int{}
	sort.Ints(nums)

	subset := []int{}

	getSubsets(nums, 0, &subset, &subsets)

	return subsets
}

func getSubsets(nums []int, i int, pSubset *[]int, pSubsets *[][]int) {
	if i >= len(nums) {
		subset := make([]int, len(*pSubset))
		copy(subset, *pSubset)
		*pSubsets = append(*pSubsets, subset)
		return
	}

	count := 1

	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
		count++
	}

	for j := 0; j <= count; j++ {

		originalLength := len(*pSubset)

		for k := 0; k < j; k++ {
			*pSubset = append(*pSubset, nums[i])
		}
		getSubsets(nums, i+1, pSubset, pSubsets)

		*pSubset = (*pSubset)[:originalLength]
	}
}
