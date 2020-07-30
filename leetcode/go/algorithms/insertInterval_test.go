package algorithms

import (
	"testing"

	"github.com/go-test/deep"
)

/*
57. Insert Interval

https://leetcode.com/problems/insert-interval/

Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
*/

func TestInsertInterval(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3},
		[]int{6, 9},
	}

	newInterval := []int{2, 5}

	expected := [][]int{
		[]int{1, 5},
		[]int{6, 9},
	}

	actual := insert(intervals, newInterval)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestInsertInterval2(t *testing.T) {
	intervals := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}

	newInterval := []int{4, 8}

	expected := [][]int{
		[]int{1, 2},
		[]int{3, 10},
		[]int{12, 16},
	}

	actual := insert(intervals, newInterval)

	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {

}
