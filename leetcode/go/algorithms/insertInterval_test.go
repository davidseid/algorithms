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
	i := 0
	currStart, currEnd := intervals[i][0], intervals[i][1]

	newStart, newEnd := newInterval[0], newInterval[1]

	if newStart < currStart && newEnd < currStart {
		result := [][]int{
			newInterval,
		}

		result = append(result, intervals[1:]...)
		return result
	}

	if newStart < currStart && newEnd <= currEnd {
		intervals[i][0] = newStart
		return intervals
	}

	if newStart < currStart && newEnd > currEnd {
		intervals[i][0] = newStart

		for i < len(intervals)-1 {
			next := intervals[i+1]

			if newEnd < next[0] {
				intervals[i][1] = newEnd
				return intervals
			}

			if newEnd >= next[0] && newEnd <= next[1] {
				intervals[i][1] = next[1]
				result := intervals[0 : i+1]
				result = append(result, intervals[i+1:]...)
				return result
			}

			if newEnd >= next[0] && newEnd > next[1] {
				newIntervals := intervals[0 : i+1]
				newIntervals = append(newIntervals, intervals[i+1:]...)
				intervals = newIntervals
				i++
			}
		}
	}

	if newStart >= currStart && newEnd <= currEnd {
		return intervals
	}

	if newStart > currStart && newEnd >= currEnd {
		for i < len(intervals)-1 {
			next := intervals[i+1]

			if newEnd < next[0] {
				intervals[i][1] = newEnd
				return intervals
			}

			if newEnd >= next[0] && newEnd <= next[1] {
				intervals[i][1] = next[1]
				result := intervals[0 : i+1]
				result = append(result, intervals[i+1:]...)
				return result
			}

			if newEnd >= next[0] && newEnd > next[1] {
				newIntervals := intervals[0 : i+1]
				newIntervals = append(newIntervals, intervals[i+1:]...)
				intervals = newIntervals
				i++
			}
		}
	}

	intervals[i][1] = newEnd

	return intervals
}
