package leetcode

import (
	"container/heap"
	"slices"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	cnt := make([]int, n)
	idle := hp{make(sort.IntSlice, n)}
	for i := range n {
		idle.IntSlice[i] = i
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	using := hp2{}

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]

		// drain
		for using.Len() > 0 && using[0].end <= start { // half-open
			heap.Push(&idle, heap.Pop(&using).(pair).i)
		}

		var i int
		if idle.Len() > 0 {
			i = heap.Pop(&idle).(int)
		} else {
			p := heap.Pop(&using).(pair)
			end += p.end - start
			i = p.i
		}

		heap.Push(&using, pair{end, i})
		cnt[i]++
	}

	return slices.Index(cnt, slices.Max(cnt))
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type pair struct{ end, i int }
type hp2 []pair

func (h hp2) Len() int { return len(h) }
func (h hp2) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.end < b.end || a.end == b.end && a.i < b.i
}
func (h hp2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
