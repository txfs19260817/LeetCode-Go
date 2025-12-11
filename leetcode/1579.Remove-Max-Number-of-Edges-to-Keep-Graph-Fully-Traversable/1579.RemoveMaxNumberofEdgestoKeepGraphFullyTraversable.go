package leetcode

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n+1) // node ranges from 1 to n
	rank := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank, n}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(p int, q int) bool { // return if actually merged
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
	uf.count--
	return true
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	aUf, bUf, used := NewUnionFind(n), NewUnionFind(n), 0
	for _, edge := range edges {
		switch edge[0] {
		case 3:
			if a, b := aUf.Union(edge[1], edge[2]), bUf.Union(edge[1], edge[2]); a || b {
				used++
			}
		}
	}
	for _, edge := range edges {
		switch edge[0] {
		case 1:
			if aUf.Union(edge[1], edge[2]) {
				used++
			}
		case 2:
			if bUf.Union(edge[1], edge[2]) {
				used++
			}
		}
	}
	if aUf.count == 1 && bUf.count == 1 {
		return len(edges) - used
	}
	return -1
}
