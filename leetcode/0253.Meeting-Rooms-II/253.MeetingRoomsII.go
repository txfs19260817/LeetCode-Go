package _253_Meeting_Rooms_II

import (
	"container/heap"
	"sort"
)

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

func minMeetingRooms(intervals [][]int) int {
	h := &intHeap{}                                                                         // Min heap stores end time
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] }) // Sort intervals by start time
	heap.Push(h, intervals[0][1])                                                           // Add the first meeting
	for _, interval := range intervals[1:] {
		if h.Len() > 0 && interval[0] >= h.Top().(int) { // If the room due to free up the earliest is free, assign that room to this meeting.
			heap.Pop(h)
		}
		// If a new room is to be assigned, then also we add to the heap.
		// If an old room is allocated, then also we have to add to the heap with updated end time.
		heap.Push(h, interval[1])
	}
	return h.Len() // The size of the heap tells us the minimum rooms required for all the meetings.
}
