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
// Use first row and first column as markers
// This requires separate variables (2) to mark whether first row and first col have zero
// Check first row and first col first and update vars
// Starting from row 2, col 2, iterate through matrix. If a zero is found
// mark the first row in col and first col in row as zero for later
// once done, for each zero in the first row, mark the col as zero
// for each zero in the first col, mark the row as zero
// zero the first row and col depending on the initial variables

// This approach uses constant space, and does not revisit any indices

// Time Complexity: O(n)
// Space Complexity: O(1)
