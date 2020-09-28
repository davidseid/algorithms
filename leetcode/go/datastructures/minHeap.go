package datastructures

type minHeap struct {
	priorityQueue []int
	numElements   int
}

func (mh *minHeap) indexOfParent(n int) int {
	if n == 1 {
		return -1
	}

	return n / 2
}

func (mh *minHeap) indexOf
