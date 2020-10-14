package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

func TestHeapSort2(t *testing.T) {
	t.Run("should sort an array of integers", func(t *testing.T) {
		input := []int{3, 7, 4, -2, 6, 0, 1}

		expected := []int{-2, 0, 1, 3, 4, 6, 7}

		actual := heapSort2(input)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}

type minHeap2 []int

func heapSort2(arr []int) []int {
	minHeap := buildHeap(arr)

	result := []int{}

	for range minHeap {
		result = append(result, minHeap.ExtractMin())
	}
	return result
}
