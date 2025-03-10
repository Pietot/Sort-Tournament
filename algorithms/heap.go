package sort

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HeapSort(array []int) []int {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range array {
		heap.Push(h, v)
	}

	sorted := make([]int, 0, len(array))
	for h.Len() > 0 {
		sorted = append(sorted, heap.Pop(h).(int))
	}

	return sorted
}
