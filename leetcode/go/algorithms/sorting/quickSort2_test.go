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
