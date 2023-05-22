package leetcode

import "container/heap"

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func connectSticks(sticks []int) int {
	ans, h := 0, intHeap(sticks)
	heap.Init(&h)
	for h.Len() > 1 {
		sum := heap.Pop(&h).(int) + heap.Pop(&h).(int)
		ans += sum
		heap.Push(&h, sum)
	}
	return ans
}
