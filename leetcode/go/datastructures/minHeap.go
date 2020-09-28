package datastructures

type minHeap []int

func (mh *minHeap) getIndexOfParent(n int) int {
	if n == 1 {
		return -1
	}

	return n / 2
}

func (mh *minHeap) getIndexOfYoungerChild(n int) int {
	return 2 * n
}

func (mh *minHeap) getIndexOfOlderChild(n int) int {
	return (2 * n) + 1
}
