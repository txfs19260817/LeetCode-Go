package leetcode

import (
	"container/heap"
	"slices"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func scheduleCourse(courses [][]int) int {
	slices.SortFunc(courses, func(a, b []int) int { return a[1] - b[1] })
	var acc int
	h := &IntHeap{}
	for _, c := range courses {
		dur, lastDay := c[0], c[1]
		if acc+dur <= lastDay {
			acc += dur
			heap.Push(h, dur)
		} else if h.Len() > 0 && dur < (*h)[0] {
			acc = acc - (*h)[0] + dur
			(*h)[0] = dur
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}
