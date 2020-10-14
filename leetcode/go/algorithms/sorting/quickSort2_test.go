package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

func TestQuickSort2(t *testing.T) {
	t.Run("should sort an array of integers", func(t *testing.T) {
		input := []int{3, 7, 4, -2, 6, 0, 1}

		expected := []int{-2, 0, 1, 3, 4, 6, 7}

		actual := quickSort2(input)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func quickSort2(arr []int) []int {
	p := partition2(arr)

	quickSort(arr[:p])

	if p < len(arr)-1 {
		quickSort(arr[p+1:])
	}
	return arr
}

func partition2(arr []int) int {
	left := 0
	right := len(arr) - 1

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	return left
}
