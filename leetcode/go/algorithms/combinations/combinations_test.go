package combinations

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

/*
77. Combinations
https://leetcode.com/problems/combinations/

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

You may return the answer in any order.

Example 1:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
Example 2:

Input: n = 1, k = 1
Output: [[1]]


Constraints:

1 <= n <= 20
1 <= k <= n


Rationale:
- Initial approach is recursive solution, each time choosing one of the next
possible choices, until we have chosen 3. It works but is inefficient because
we are unnecessarily copying the current slice into a next slice on each iteration,
and additionally copying the remaining choices into a new set of choices.
This is too expensive for time and memory.

Optimization ideas:
Instead of using the actual array of choices and actual array of curr combination,
could instead maintain array of the indexes to turn on, once we've reached 3, use them to grab the
relevant numbers and add to the combinations
*/

func combine(n int, k int) [][]int {
	combinations := [][]int{}

	choices := []int{}

	for i := 1; i <= n; i++ {
		choices = append(choices, i)
	}

	curr := []int{}

	getCombinations(curr, choices, k, &combinations)
	fmt.Println(combinations)
	return combinations
}

func getCombinations(curr []int, choices []int, remaining int, combinations *[][]int) {

	if remaining == 0 {
		*combinations = append(*combinations, curr)
		return
	}

	for i, v := range choices {

		next := make([]int, len(curr))
		copy(next, curr)
		next = append(next, v)

		nextChoices := []int{}

		for j := i + 1; j < len(choices); j++ {
			nextChoices = append(nextChoices, choices[j])
		}

		getCombinations(next, nextChoices, remaining-1, combinations)
	}
}

func TestCombine(t *testing.T) {
	t.Run("should return all possible combinations of k numbers from 1..n", func(t *testing.T) {
		expected := [][]int{
			{2, 4},
			{3, 4},
			{2, 3},
			{1, 2},
			{1, 3},
			{1, 4},
		}

		actual := combine(4, 2)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should return all possible combinations of k numbers from 1..n when n and k are 1", func(t *testing.T) {
		expected := [][]int{
			{1},
		}

		actual := combine(1, 1)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
