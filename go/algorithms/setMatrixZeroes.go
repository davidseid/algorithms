package algorithms

// https://leetcode.com/problems/set-matrix-zeroes/
// Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
// Solve in place
// Then, solve with constant space

func SetZeroes(matrix [][]int) {
	for _, row := range matrix {
		if hasZero(row) {
			zeroRow(row)
		}
	}
}

func hasZero(row []int) bool {
	for _, value := range row {
		if value == 0 {
			return true
		}
	}
	return false
}

func zeroRow(row []int) {
	for col := range row {
		row[col] = 0
	}
}

// Approach:
// For each row
// if rowHasZero
// zeroRow()

// Time Complexity: O(n)
// Space Complexity: O(1)
