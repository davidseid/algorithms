package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

/*
Merge sort
*/

func mergeSort(arr []int) {
	if len(arr) == 1 {
		return
	}

	mid := len(arr) / 2
	front := arr[:mid]
	back := arr[mid+1:]

	mergeSort(front)
	mergeSort(back)

	merge(front, back)
}

func merge(front []int, back []int) []int {
	result := []int{}

	fi := 0
	bi := 0

	for fi < len(front) || bi < len(back) {
		if fi < len(front) && front[fi] <= back[bi] {
			result = append(result, front[fi])
			fi++
		} else {
			result = append(result, back[bi])
			bi++
		}
	}
	return result
}

func TestMergeSort(t *testing.T) {
	t.Run("should sort an array without dupes", func(t *testing.T) {
		input := []int{5, 3, 1, -3, 9, 6, 7}

		mergeSort(input)

		expected := []int{-3, 1, 3, 5, 6, 7, 9}

		if diff := deep.Equal(input, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should sort an array with dupes", func(t *testing.T) {
		input := []int{5, 3, 9, 1, -3, 9, 6, 7}

		mergeSort(input)

		expected := []int{-3, 1, 3, 5, 6, 7, 9, 9}

		if diff := deep.Equal(input, expected); diff != nil {
			t.Error(diff)
		}
	})
}
