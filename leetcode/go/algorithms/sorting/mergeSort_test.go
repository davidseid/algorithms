package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

/*
Merge sort
*/

func mergeSort(arr []int) {

}

func merge(arr1 []int, arr2 []int) []int {

}

func TestMergeSort(t *testing.T) {
	t.Run("should sort an array without dupes", func(t *testing.T) {
		input := []int{5, 3, 1, -3, 9, 6, 7}

		actual := mergeSort(input)

		expected := []int{-3, 1, 3, 5, 6, 7, 9}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
