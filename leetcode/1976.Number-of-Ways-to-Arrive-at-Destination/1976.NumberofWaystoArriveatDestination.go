package leetcode

import (
	"container/heap"
	"math"
)

const mod = 1_000_000_007

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

func countPaths(n int, roads [][]int) int {
	ways, dist := make([]int, n), make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	ways[0], dist[0] = 1, 0

	g := map[int][]pair{}
	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		g[u] = append(g[u], pair{v, w})
		g[v] = append(g[v], pair{u, w})
	}

	h := &pairHeap{pair{v: 0, w: 0}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		u, w := cur.v, cur.w
		if w > dist[u] {
			continue
		}
		for _, neighbor := range g[u] {
			nextW := w + neighbor.w
			if nextW < dist[neighbor.v] {
				dist[neighbor.v], ways[neighbor.v] = nextW, ways[u]
				heap.Push(h, pair{v: neighbor.v, w: nextW})
			} else if nextW == dist[neighbor.v] {
				ways[neighbor.v] = (ways[neighbor.v] + ways[u]) % mod
			}
		}
	}

	return ways[n-1]
}
