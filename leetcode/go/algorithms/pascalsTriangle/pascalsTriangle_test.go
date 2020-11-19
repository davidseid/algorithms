package pascalsTriangle

import "testing"

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

		if actual != expected {
			t.Errorf("Got %d, want %d", actual, expected)
		}
	})

}
