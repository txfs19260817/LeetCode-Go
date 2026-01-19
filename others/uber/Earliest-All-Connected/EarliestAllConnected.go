package uber

import "sort"

// RiderLog represents a shared-ride connection between two riders at a given time.
type RiderLog struct {
	Time int
	A    int
	B    int
}

// EarliestAllConnected returns the earliest time when all riders are connected.
// Riders are labeled from 0 to n-1. If never fully connected, returns -1.
func EarliestAllConnected(n int, logs []RiderLog) int {
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return 0
	}

	sort.Slice(logs, func(i, j int) bool { return logs[i].Time < logs[j].Time })

	uf := newUnionFind(n)
	for _, log := range logs {
		if log.A < 0 || log.A >= n || log.B < 0 || log.B >= n {
			continue
		}
		if uf.union(log.A, log.B) && uf.count == 1 {
			return log.Time
		}
	}
	return -1
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{parent: parent, rank: rank, count: n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(a, b int) bool {
	ra := uf.find(a)
	rb := uf.find(b)
	if ra == rb {
		return false
	}
	if uf.rank[ra] < uf.rank[rb] {
		uf.parent[ra] = rb
	} else if uf.rank[ra] > uf.rank[rb] {
		uf.parent[rb] = ra
	} else {
		uf.parent[rb] = ra
		uf.rank[ra]++
	}
	uf.count--
	return true
}
