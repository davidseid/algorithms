package datastructures

type minHeap []int

func (mh *minHeap) indexOfParent(n int) int {
	if n == 1 {
		return -1
	}

	return n / 2
}

func (mh *minHeap) indexOfYoungerChild(n int) int {
	return 2 * n
}

func (mh *minHeap) indexOfOlderChild(n int) int {
	return (2 * n) + 1
}
