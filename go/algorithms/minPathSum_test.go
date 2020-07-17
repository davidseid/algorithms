package algorithms

import (
	"testing"
)

/*
64. Minimum Path Sum
https://leetcode.com/problems/minimum-path-sum/

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

func minPathSum(grid [][]int) int {

	min := getMinPathSum(grid, 0, 0, grid[0][0], -1)
	return min
}

func getMinPathSum(grid [][]int, row int, col int, sum int, min int) int {

	if row+1 >= len(grid) && col+1 >= len(grid[0]) {
		if min < 0 || sum < min {
			return sum
		} else {
			return min
		}
	}

	var min1 int
	var min2 int

	if row+1 < len(grid) {
		min1 = getMinPathSum(grid, row+1, col, sum+grid[row+1][col], min)
	}

	if col+1 < len(grid[0]) {
		min2 = getMinPathSum(grid, row, col+1, sum+grid[row][col+1], min)
	}

	if min1 < min2 {
		return min1
	}
	return min2
}

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	expected := 7

	actual := minPathSum(grid)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
