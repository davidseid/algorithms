package algorithms

import "fmt"

// https://leetcode.com/problems/set-matrix-zeroes/
// Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
// Solve in place
// Then, solve with constant space

func SetZeroes(matrix [][]int) {
	var shouldZeroFirstRow bool
	var shouldZeroFirstCol bool

	if firstRowHasZeroes(matrix) {
		shouldZeroFirstRow = true
	}

	if firstColHasZeroes(matrix) {
		shouldZeroFirstCol = true
	}

	markMatrix(matrix)
	fmt.Println("MARKED MATRIX")
	fmt.Println(matrix)

	zeroColumns(matrix)
	zeroRows(matrix)

	if shouldZeroFirstRow {
		zeroFirstRow(matrix)
	}

	if shouldZeroFirstCol {
		zeroFirstCol(matrix)
	}
}

func firstRowHasZeroes(matrix [][]int) bool {
	for col := range matrix[0] {
		if matrix[0][col] == 0 {
			return true
		}
	}
	return false
}

func firstColHasZeroes(matrix [][]int) bool {
	for row := range matrix {
		if matrix[row][0] == 0 {
			return true
		}
	}
	return false
}

func markMatrix(matrix [][]int) {
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				matrix[0][col] = 0
				matrix[row][0] = 0
			}
		}
	}
}

func zeroColumns(matrix [][]int) {
	for col := 1; col < len(matrix[0]); col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < len(matrix); row++ {
				matrix[row][col] = 0
			}
		}
	}
}

func zeroRows(matrix [][]int) {
	for row := 1; row < len(matrix); row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < len(matrix[row]); col++ {
				matrix[row][col] = 0
			}
		}
	}
}

func zeroFirstRow(matrix [][]int) {
	for col := range matrix[0] {
		matrix[0][col] = 0
	}
}

func zeroFirstCol(matrix [][]int) {
	for row := range matrix {
		matrix[row][0] = 0
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
