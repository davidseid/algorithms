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

func (mh *minHeap2) Insert(val int) {
	*mh = append(*mh, val)
	mh.bubbleUp(len(*mh) - 1)
}

func (mh *minHeap2) bubbleUp(i int) {
	parentIndex := mh.getParent(i)
	if (*mh)[i] < (*mh)[parentIndex] {
		mh.swap(i, parentIndex)
	}
	mh.bubbleUp(parentIndex)
}

func (mh *minHeap2) swap(a, b int) {
	(*mh)[a], (*mh)[b] = (*mh)[b], (*mh)[a]
}

func (mh *minHeap2) getParent(i int) int {
	return i / 2
}

func buildHeap(arr []int) minHeap2 {
	heap := minHeap2{0}

	for _, v := range arr {
		heap.Insert(v)
	}
}

func heapSort2(arr []int) []int {
	minHeap := buildHeap(arr)

	result := []int{}

	for range minHeap {
		result = append(result, minHeap.ExtractMin())
	}
	return result
}
