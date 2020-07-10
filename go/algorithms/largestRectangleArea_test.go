package algorithms

import "testing"

/*
Source: https://leetcode.com/problems/largest-rectangle-in-histogram/

Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

The largest rectangle is shown in the shaded area, which has area = 10 unit.

Example:

Input: [2,1,5,6,2,3]
Output: 10
*/

func TestLargestRectangleArea(t *testing.T) {
	histogram := []int{2, 1, 5, 6, 2, 3}

	expected := 10
	actual := largestRectangleArea(histogram)

	if actual != expected {
		t.Errorf("got %d, want %d", actual, expected)
	}
}
