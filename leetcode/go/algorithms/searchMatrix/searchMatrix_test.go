package searchMatrix

import "testing"

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

	row := findRow(&matrix, 0, n, target)

	return findCol(row, target)
}

func findRow(pMatrix *[][]int, start int, end int, target int) []int {
	matrix := *pMatrix

	m := end - start

	if m == 1 {
		return matrix[0]
	}

	pivot := m / 2

	if target >= matrix[pivot][0] {
		return findRow(pMatrix, pivot, end, target)
	}

	return findRow(pMatrix, start, pivot, target)

}

func findCol(row []int, target int) bool {
	n := len(row)

	if n == 1 {
		if row[0] == target {
			return true
		}

		return false
	}

	pivot := n / 2

	if target >= row[pivot] {
		return findCol(row[pivot:], target)
	}

	return findCol(row[:pivot], target)
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
