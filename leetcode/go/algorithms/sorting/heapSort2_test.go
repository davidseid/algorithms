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

func (mh *minHeap2) getChildren(i int) (left, right int) {
	return i * 2, i*2 + 1
}

func (mh *minHeap2) ExtractMin() int {
	// swap root and end
	// grab end
	// bubbledown root
	if len(*mh) < 2 {
		panic("Nothing left in the heap")
	}

	mh.swap(1, len(*mh)-1)
	min := (*mh)[len(*mh)-1]
	(*mh) = (*mh)[:len(*mh)-1]

	mh.bubbleDown(1)
	return min
}

func (mh *minHeap2) bubbleDown(i int) {
	target := i
	left, right := mh.getChildren(i)

	if (*mh)[i] > (*mh)[left] {
		target = left
	}

	if (*mh)[i] > (*mh)[right] {
		target = right
	}

	mh.swap(i, target)
	mh.bubbleDown(target)
}

func buildHeap(arr []int) minHeap2 {
	heap := minHeap2{0}

	for _, v := range arr {
		heap.Insert(v)
	}
	return heap
}

func heapSort2(arr []int) []int {
	minHeap := buildHeap(arr)

	result := []int{}

	for range minHeap {
		result = append(result, minHeap.ExtractMin())
	}
	return result
}
