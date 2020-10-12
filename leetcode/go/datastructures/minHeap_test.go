package datastructures

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

type minHeap []int

func (mh *minHeap) getParent(i int) int {
	if i == 1 {
		return -1
	}

	return i / 2
}

func (mh *minHeap) getLeftChild(i int) int {
	return 2 * i
}

func (mh *minHeap) getRightChild(i int) int {
	return (2 * i) + 1
}

func (mh *minHeap) insert(item int) {
	*mh = append(*mh, item)
	mh.bubbleUp(len(*mh) - 1)
}

func (mh *minHeap) bubbleUp(i int) {
	if mh.getParent(i) == -1 {
		return
	}

	if (*mh)[mh.getParent(i)] > (*mh)[i] {
		mh.swap(i, mh.getParent(i))
		mh.bubbleUp(mh.getParent(i))
	}
}

func (mh *minHeap) swap(a, b int) {
	(*mh)[a], (*mh)[b] = (*mh)[b], (*mh)[a]
}

func makeHeap(items []int) minHeap {
	heap := minHeap{0}
	for _, v := range items {
		fmt.Println(heap)
		heap.insert(v)
	}
	return heap
}

func (mh *minHeap) extractMin() int {
	// swap root and bottom right
	// pop off the bottom right
	// bubble down the root

	if len(*mh) < 2 {
		return -1
	}

	min := (*mh)[1]
	mh.swap(1, len(*mh)-1)
	*mh = (*mh)[:len(*mh)-1]
	mh.bubbleDown(1)
	return min
}

func (mh *minHeap) bubbleDown(i int) {
	target := i
	left := mh.getLeftChild(i)
	right := mh.getRightChild(i)

	if (*mh)[target] > (*mh)[left] {
		target = (*mh)[left]
	}

	if (*mh)[target] > (*mh)[right] {
		target = (*mh)[right]
	}

	if target != i {
		mh.swap(i, target)
		mh.bubbleDown(target)
	}
}

func TestMinHeap(t *testing.T) {
	t.Run("should build a min heap", func(t *testing.T) {
		input := []int{1, 5, 6, 8, 9, 7, 3}

		actual := makeHeap(input)

		expected := minHeap{0, 1, 5, 3, 8, 9, 7, 6}

		if diff := deep.Equal(actual, expected); diff != nil {
			fmt.Println(actual, expected)
			t.Error(diff)
		}
	})
}
