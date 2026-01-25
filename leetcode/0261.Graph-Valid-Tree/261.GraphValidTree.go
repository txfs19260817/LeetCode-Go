package leetcode

type UnionFind struct {
	parent, rank []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, make([]int, n)}
}

func (uf *UnionFind) Find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(p, q int) bool {
	p, q = uf.Find(p), uf.Find(q)
	if p == q {
		return false
	}
	if uf.rank[p] < uf.rank[q] {
		uf.parent[p] = q
	} else {
		uf.parent[q] = p
		if uf.rank[p] == uf.rank[q] {
			uf.rank[p]++
		}
	}
	return true
}

func validTree(n int, edges [][]int) bool {
	// A tree must have exactly n-1 edges.
	// If it has fewer, it's disconnected.
	// If it has more, it has a cycle.
	if len(edges) != n-1 {
		return false
	}

	uf := NewUnionFind(n)
	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			// Cycle detected
			return false
		}
	}
	// If we processed n-1 edges without a cycle, the graph is a tree.
	// (n nodes, n-1 edges, no cycles implies connected)
	return true
}
