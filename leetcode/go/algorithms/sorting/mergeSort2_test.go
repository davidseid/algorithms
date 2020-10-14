package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

func TestMergeSort2(t *testing.T) {
	t.Run("should sort an array of integers", func(t *testing.T) {
		input := []int{3, 7, 4, -2, 6, 0, 1}

		expected := []int{-2, 0, 1, 3, 4, 6, 7}

		actual := mergeSort2(input)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func mergeSort2(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	return merge2(mergeSort2(arr[:mid]), mergeSort2(arr[mid:]))
}

func merge2(left, right []int) []int {
	result := []int{}

	l := 0
	r := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
			continue
		}
		result = append(result, right[r])
		r++
	}

	for l < len(left) {
		result = append(result, left[l])
		l++
	}

	for r < len(right) {
		result = append(result, right[r])
		r++
	}

	return result
}
