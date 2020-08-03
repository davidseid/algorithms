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

// Dynamic Programming Solution
func minPathSum(grid [][]int) int {

	dp := copyGrid(grid)

	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	dp[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j-1] < dp[i-1][j] {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func copyGrid(grid [][]int) [][]int {
	copy := [][]int{}

	for range grid {
		row := []int{}
		for range grid[0] {
			row = append(row, -1)
		}
		copy = append(copy, row)
	}

	return copy
}

// Recursive Solution with memoization
// func minPathSum(grid [][]int) int {
// 	min := -1
// 	memo := map[string]int{}

// 	getMinPathSum(grid, 0, 0, grid[0][0], &min, memo)
// 	return min
// }

// func getMinPathSum(grid [][]int, row int, col int, sum int, min *int, memo map[string]int) {
// 	key := fmt.Sprintf("%d,%d", row, col)

// 	if pathMin, ok := memo[key]; ok {
// 		*min = pathMin
// 	}

// 	if row+1 >= len(grid) && col+1 >= len(grid[0]) {
// 		if *min < 0 || sum < *min {
// 			*min = sum
// 		}
// 	}

// 	if row+1 < len(grid) {
// 		getMinPathSum(grid, row+1, col, sum+grid[row+1][col], min, memo)
// 	}

// 	if col+1 < len(grid[0]) {
// 		getMinPathSum(grid, row, col+1, sum+grid[row][col+1], min, memo)
// 	}
// 	memo[key] = *min
// }

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
