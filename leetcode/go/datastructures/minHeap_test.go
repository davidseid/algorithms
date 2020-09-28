package datastructures

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

type minHeap []int

func (mh *minHeap) getParent(n int) int {
	if n == 1 {
		return -1
	}

	return n / 2
}

func (mh *minHeap) getYoungerChild(n int) int {
	return 2 * n
}

func (mh *minHeap) getOlderChild(n int) int {
	return (2 * n) + 1
}

func (mh *minHeap) insert(item int) {
	*mh = append(*mh, item)
	mh.bubbleUp(len(*mh) - 1)
}

func (mh *minHeap) bubbleUp(p int) {
	if mh.getParent(p) == -1 {
		return
	}

	if (*mh)[mh.getParent(p)] > (*mh)[p] {
		mh.swap(p, mh.getParent(p))
		mh.bubbleUp(mh.getParent(p))
	}
}

func (mh *minHeap) swap(a, b int) {
	(*mh)[a], (*mh)[b] = (*mh)[b], (*mh)[a]
}

func makeHeap(items []int) minHeap {
	var heap minHeap
	for _, v := range items {
		heap.insert(v)
	}
	return heap
}

func TestMinHeap(t *testing.T) {
	t.Run("should build a min heap", func(t *testing.T) {
		input := []int{5, 3, 7, 9, 2, 31, 19, 2}

		actual := makeHeap(input)

		expected := minHeap{5, 2, 3, 2, 7, 31, 19, 9}

		if diff := deep.Equal(actual, expected); diff != nil {
			fmt.Println(actual, expected)
			t.Error(diff)
		}

	})
}
