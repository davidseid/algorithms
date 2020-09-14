package sorting

import (
	"testing"

	"github.com/go-test/deep"
)

/*
Merge sort
*/

func mergeSort(arr []int, low int, high int) {
	var mid int

	if low < high {
		mid = (low + high) / 2
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}

}

func merge(arr []int, low int, mid int, high int) {
	fi := 0
	front := []int{}
	bi := 0
	back := []int{}

	for i := low; i <= mid; i++ {
		front = append(front, arr[i])
	}

	for i := mid + 1; i <= high; i++ {
		back = append(back, arr[i])
	}

	i := low

	for fi < len(front) && bi < len(back) {
		if front[fi] <= back[bi] {
			arr[i] = front[fi]
			fi++
		} else {
			arr[i] = back[bi]
			bi++
		}
		i++
	}

	for fi < len(front) {
		arr[i] = front[fi]
		fi++
		i++
	}

	for bi < len(back) {
		arr[i] = back[bi]
		bi++
		i++
	}
}

func TestMergeSort(t *testing.T) {
	t.Run("should sort an array without dupes", func(t *testing.T) {
		input := []int{5, 3, 1, -3, 9, 6, 7}

		mergeSort(input, 0, len(input)-1)

		expected := []int{-3, 1, 3, 5, 6, 7, 9}

		if diff := deep.Equal(input, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should sort an array with dupes", func(t *testing.T) {
		input := []int{5, 3, 9, 1, -3, 9, 6, 7}

		mergeSort(input, 0, len(input)-1)

		expected := []int{-3, 1, 3, 5, 6, 7, 9, 9}

		if diff := deep.Equal(input, expected); diff != nil {
			t.Error(diff)
		}
	})
}
