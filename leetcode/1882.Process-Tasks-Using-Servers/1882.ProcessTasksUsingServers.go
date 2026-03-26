package leetcode

import "container/heap"

type tuple struct {
	t, w, idx int
}

type tupleHeap struct {
	data []tuple
	busy bool // false: idle堆, true: busy堆
}

func (h tupleHeap) Len() int { return len(h.data) }

func (h tupleHeap) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]

	if h.busy {
		// busy: 按 (t, w, idx)
		if a.t != b.t {
			return a.t < b.t
		}
		if a.w != b.w {
			return a.w < b.w
		}
		return a.idx < b.idx
	}

	// idle: 按 (w, idx)
	if a.w != b.w {
		return a.w < b.w
	}
	return a.idx < b.idx
}

func (h tupleHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *tupleHeap) Push(x interface{}) {
	h.data = append(h.data, x.(tuple))
}

func (h *tupleHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}

func assignTasks(servers []int, tasks []int) []int {
	ans := make([]int, len(tasks))
	idle, using := &tupleHeap{}, &tupleHeap{busy: true}
	for i, w := range servers {
		heap.Push(idle, tuple{w: w, idx: i})
	}
	for i, t := range tasks {
		for using.Len() > 0 && using.data[0].t <= i {
			top := heap.Pop(using).(tuple)
			heap.Push(idle, tuple{w: top.w, idx: top.idx})
		}
		if idle.Len() > 0 {
			top := heap.Pop(idle).(tuple)
			ans[i] = top.idx
			heap.Push(using, tuple{t: t + i, w: top.w, idx: top.idx})
		} else {
			top := heap.Pop(using).(tuple)
			ans[i] = top.idx
			heap.Push(using, tuple{t: top.t + t, w: top.w, idx: top.idx})
		}
	}
	return ans
}
