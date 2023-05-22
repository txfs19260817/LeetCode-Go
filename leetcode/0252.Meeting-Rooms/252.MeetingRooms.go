package leetcode

import (
	"container/heap"
	"sort"
)

type intervalHeap [][]int

func (h intervalHeap) Len() int {
	return len(h)
}

func (h intervalHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h intervalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intervalHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *intervalHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) < 2 {
		return true
	}
	var h intervalHeap = intervals
	heap.Init(&h)
	for pre := heap.Pop(&h).([]int); h.Len() > 0; {
		cur := heap.Pop(&h).([]int)
		if cur[0] < pre[1] {
			return false
		}
		pre = cur
	}
	return true
}

func canAttendMeetings2(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] < intervals[i][1] {
			return false
		}
	}
	return true
}
