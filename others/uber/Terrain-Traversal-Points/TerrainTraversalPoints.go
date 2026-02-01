package uber

import (
	"container/heap"
	"sort"
)

type Cell struct {
	val, r, c int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Query struct {
	limit, idx int
}

// MaxPoints computes the maximum points achievable for each limit.
func MaxPoints(terrain [][]int, limits []int) []int {
	m, n := len(terrain), len(terrain[0])
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	queries := make([]Query, len(limits))
	for i, v := range limits {
		queries[i] = Query{v, i}
	}

	sort.Slice(queries, func(i, j int) bool { return queries[i].limit < queries[j].limit })

	ans := make([]int, len(limits))
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := &MinHeap{}

	// Start from (0,0)
	// We push (0,0) to heap initially.
	// We mark it visited immediately so it's not added again.
	heap.Push(h, Cell{terrain[0][0], 0, 0})
	visited[0][0] = true

	var count int
	for _, q := range queries {
		for h.Len() > 0 {
			// Check if the minimum value reachable is passable with current limit
			minCell := (*h)[0]
			if minCell.val >= q.limit {
				break
			}

			// Pop and count
			curr := heap.Pop(h).(Cell)
			count++

			// Add neighbors
			for _, d := range dirs {
				nr, nc := curr.r+d[0], curr.c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
					visited[nr][nc] = true
					heap.Push(h, Cell{terrain[nr][nc], nr, nc})
				}
			}
		}
		ans[q.idx] = count
	}

	return ans
}
