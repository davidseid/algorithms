package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

/*
Quicksort

Rationale:
Divide and Conquer recursive algorithm.
Choose a pivot point (can be anything, random is best on average).
Swap the elements so that all the elements less than the pivot element are left
and the greater are on the write (the two sides need not be sorted themselves)

Then call quick sort on each side of the pivot.

Key Insight:
The pivot itself will be moved due to swapping. It is easiest to start with
or move the pivot element to the end, then we can keep a pointer for where the left side is
that will eventually be swapped with the right. This left pointer is the first element in the array
that is higher than the pivot element.

At the end of the iteration, move the pivot element to that index and return that index as the parition point

Time Complexity: O(nlog(n)) in the average case. This is the case because on average the pivot index will roughly
evenly divide the array, resulting in logn recursive calls, each of which require an iteration loop. It can be worse O(n^2)
if we consistently pick the worst pivot point.
*/

func quickSort(arr []int) {
	if len(arr) > 0 {
		p := partition(arr)
		quickSort(arr[:p])

		if p < len(arr)-1 {
			quickSort(arr[p+1:])
		}
	}
}

func partition(arr []int) int {
	left := 0
	right := len(arr) - 1

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	return left
}

func TestQuickSort(t *testing.T) {
	t.Run("should sort an array without dupes", func(t *testing.T) {
		input := []int{5, 3, 1, -3, 9, 6, 7}

		quickSort(input)

		expected := []int{-3, 1, 3, 5, 6, 7, 9}

		if diff := deep.Equal(input, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should sort an array with dupes", func(t *testing.T) {
		input := []int{5, 3, 9, 1, -3, 9, 6, 7}

		quickSort(input)

		expected := []int{-3, 1, 3, 5, 6, 7, 9, 9}

		if diff := deep.Equal(input, expected); diff != nil {
			t.Error(diff)
		}
	})
}
