package leetcode

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	type cp struct {
		c, p int
	}
	cps := make([]cp, len(profits))
	for i := range profits {
		cps[i] = cp{c: capital[i], p: profits[i]}
	}
	sort.Slice(cps, func(i, j int) bool {
		return cps[i].c < cps[j].c
	})
	var h intHeap
	for i := 0; k > 0; k-- {
		for ; i < len(cps) && w >= cps[i].c; i++ {
			heap.Push(&h, cps[i].p)
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(&h).(int)
	}
	return w
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
