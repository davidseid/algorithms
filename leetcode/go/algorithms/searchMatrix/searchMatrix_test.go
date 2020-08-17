package searchMatrix

import (
	"testing"
)

/*
74. Search a 2D Matrix
https://leetcode.com/problems/search-a-2d-matrix/

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false

Rationale:
Use binary search on the matrix, twice. Use binary search to find the row where the target may be.
Use binary search on the row to find the target if it exists.

Optimizations:
Initial implementation copied the matrix slices, the optimization used a pointer to a matrix
and used start and end arguments as opposed to capturing slices. Surprisingly, this took more memory
not less. It appears golang slices are already optimized for this usage, since they are already
pointing at an underlying array, they are themselves just managing the indices pointing to subsets of the array.

Time Complexity:
O(logn) + O(logm), since we need to use binary search on both m and then n

Space Complexity:
O(logn) + O(logm) since we make a recursive call on each part of the binary search.
*/

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)

	if m == 0 {
		return false
	}

	n := len(matrix[0])

	if n == 0 {
		return false
	}

	row := findRow(&matrix, 0, m, target)

	return findCol(&matrix, row, 0, n, target)
}

func findRow(pMatrix *[][]int, start int, end int, target int) int {
	matrix := *pMatrix

	m := end - start

	if m == 1 {
		return start
	}

	pivot := m/2 + start

	if target >= matrix[pivot][0] {
		return findRow(pMatrix, pivot, end, target)
	}

	return findRow(pMatrix, start, pivot, target)

}

func findCol(pMatrix *[][]int, row int, start int, end int, target int) bool {
	matrix := *pMatrix
	n := end - start

	if n == 1 {
		if matrix[row][start] == target {
			return true
		}

		return false
	}

	pivot := n/2 + start

	if target >= matrix[row][pivot] {
		return findCol(pMatrix, row, pivot, end, target)
	}

	return findCol(pMatrix, row, start, pivot, target)
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	target := 3

	actual := searchMatrix(matrix, target)
	expected := true

	if actual != expected {
		t.Errorf("Got %v, want %v", actual, expected)
	}
}

func TestSearchMatrix2(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	target := 13

	actual := searchMatrix(matrix, target)
	expected := false

	if actual != expected {
		t.Errorf("Got %v, want %v", actual, expected)
	}
}
