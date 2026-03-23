package leetcode

import (
	"container/heap"
	"math"
	"slices"
)

type pair struct {
	v, w int
}

type pairHeap []pair

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Less(i, j int) bool  { return h[i].w < h[j].w }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[k-1] = 0

	g := map[int][]pair{}
	for _, t := range times {
		u, v, w := t[0]-1, t[1]-1, t[2]
		g[u] = append(g[u], pair{v, w})
	}

	h := &pairHeap{{k - 1, 0}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		u, w := p.v, p.w
		if w > dist[u] {
			continue
		}
		for _, neighbor := range g[u] {
			v, nextW := neighbor.v, neighbor.w
			if w+nextW < dist[v] {
				dist[v] = w + nextW
				heap.Push(h, pair{v, dist[v]})
			}
		}
	}
	if ans := slices.Max(dist); ans < math.MaxInt {
		return ans
	}
	return -1
}
