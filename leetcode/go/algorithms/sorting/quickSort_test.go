package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

/*
Quicksort

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
