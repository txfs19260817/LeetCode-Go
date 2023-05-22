package leetcode

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	var ans int
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	h := &intHeap{}
	for time := 1; time <= 1e5; time++ {
		for len(events) > 0 && time == events[0][0] {
			heap.Push(h, events[0][1])
			events = events[1:]
		}
		for h.Len() > 0 && time > h.Top().(int) {
			heap.Pop(h)
		}
		if h.Len() > 0 && time <= heap.Pop(h).(int) { // attend this popped meeting at current `time`
			ans++
		}
	}
	return ans
}

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

func (h *intHeap) Top() interface{} {
	return (*h)[0]
}
