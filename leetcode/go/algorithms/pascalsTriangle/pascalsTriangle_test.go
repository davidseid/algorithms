package pascalsTriangle

import (
	"testing"

	"github.com/go-test/deep"
)

/*
119. Pascal's Triangle II

https://leetcode.com/problems/pascals-triangle-ii/

Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.

Notice that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Follow up:

Could you optimize your algorithm to use only O(k) extra space?

*/

func TestGetRowFromPascalsTriangle(t *testing.T) {

	t.Run("should work", func(t *testing.T) {
		rowIndex := 3
		expected := []int{1, 3, 3, 1}

		actual := getRow(rowIndex)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should work2", func(t *testing.T) {
		rowIndex := 0
		expected := []int{1}

		actual := getRow(rowIndex)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should work3", func(t *testing.T) {
		rowIndex := 1
		expected := []int{1, 1}

		actual := getRow(rowIndex)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})

}

func getRow(rowIndex int) []int {
	triangle := [][]int{
		{1},
		{1, 1},
	}

	for i := 2; i <= rowIndex; i++ {
		row := []int{1}

		for j := 1; j < i; j++ {
			left := triangle[i-1][j-1]
			right := triangle[i-1][j]
			row = append(row, left+right)
		}

		row = append(row, 1)
		triangle = append(triangle, row)
	}

	return triangle[rowIndex]
}