package datastructures

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

/*
Lessons learned from implementation:
HeapSort itself is simple to implement once a heap is built,
simply iterate through the range of numbers and repeatedly extract the minimum.
Since extract minimum is a O(logn) operation, heapsort is O(nlogn)
*/

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

	if left < len(*mh) && (*mh)[target] > (*mh)[left] {
		target = left
	}

	if right < len(*mh) && (*mh)[target] > (*mh)[right] {
		target = right
	}

	if target != i {
		mh.swap(i, target)
		mh.bubbleDown(target)
	}
}

func heapSort(items []int) []int {
	heap := makeHeap(items)

	for i := range items {
		items[i] = heap.extractMin()
	}
	return items
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

	t.Run("should extract minimimum from a min heap", func(t *testing.T) {
		heap := makeHeap([]int{1, 5, 6, 8, 9, 7, 3})

		expected := 1
		actual := heap.extractMin()

		if actual != expected {
			t.Errorf("Got %d, expected %d", actual, expected)
		}
	})

	t.Run("should enable heap sorting", func(t *testing.T) {
		actual := heapSort([]int{1, 5, 6, 8, 9, 7, 3})
		expected := []int{1, 3, 5, 6, 7, 8, 9}

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
