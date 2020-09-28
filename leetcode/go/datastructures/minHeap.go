package datastructures

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
