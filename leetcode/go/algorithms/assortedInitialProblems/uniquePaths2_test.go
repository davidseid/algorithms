package algorithms

import (
	"testing"
)

/*
63. Unique Paths II
https://leetcode.com/problems/unique-paths-ii/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

func TestRecursiveUniquePaths2(t *testing.T) {
	input := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}

	expected := 2

	actual := recursiveUniquePathsWithObstacles(input)

	if actual != expected {
		t.Errorf("Got %d, wanted %d", actual, expected)
	}
}

func TestDynamicProgrammingUniquePaths2(t *testing.T) {
	input := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}

	expected := 2

	actual := uniquePathsWithObstacles(input)

	if actual != expected {
		t.Errorf("Got %d, wanted %d", actual, expected)
	}
}

// Dynamic Programming Solution
// Time Complexity: O(m*n), we must look at each cell once
// Space Complexity: O(m*n), we must make a copy of the grid
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)

	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := [][]int{}

	for i := 0; i < m; i++ {
		row := []int{}

		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		dp = append(dp, row)
	}

	dp[m-1][n-1] = 1

	for col := n - 2; col >= 0; col-- {
		if obstacleGrid[m-1][col] == 1 {
			dp[m-1][col] = 0
		} else {
			dp[m-1][col] = dp[m-1][col+1]
		}
	}

	for row := m - 2; row >= 0; row-- {
		if obstacleGrid[row][n-1] == 1 {
			dp[row][n-1] = 0
		} else {
			dp[row][n-1] = dp[row+1][n-1]
		}
	}

	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			if obstacleGrid[row][col] == 1 {
				dp[row][col] = 0
			} else {
				rightPaths := dp[row][col+1]
				downPaths := dp[row+1][col]

				dp[row][col] = rightPaths + downPaths
			}
		}
	}

	return dp[0][0]
}

// Recursive Solution
// Time complexity: O(2^mm) For each cell, we may branch in two directions
// Space complexity: O(2^mn) Each recursive call increases the call stack
func recursiveUniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)

	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	return countUniquePaths(&obstacleGrid, 0, 0, m, n)
}

func countUniquePaths(pGrid *[][]int, row int, col int, m int, n int) int {
	grid := *pGrid

	if row == m-1 && col == n-1 && grid[row][col] != 1 {
		return 1
	}

	paths := 0

	if row+1 < m && grid[row+1][col] != 1 {
		paths += countUniquePaths(pGrid, row+1, col, m, n)
	}

	if col+1 < n && grid[row][col+1] != 1 {
		paths += countUniquePaths(pGrid, row, col+1, m, n)
	}

	return paths
}

// func printBoard(grid [][]int) {
// 	for _, row := range grid {
// 		fmt.Println(row)
// 	}
// 	fmt.Println("------")
// }
