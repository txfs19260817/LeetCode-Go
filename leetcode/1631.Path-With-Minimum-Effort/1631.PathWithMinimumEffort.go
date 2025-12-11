package leetcode

import "sort"

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(p int, q int) {
	p, q = uf.Find(p), uf.Find(q)
	if p == q {
		return
	}
	if uf.rank[p] < uf.rank[q] {
		uf.parent[p] = q
	} else {
		uf.parent[q] = p
		if uf.rank[p] == uf.rank[q] {
			uf.rank[p]++
		}
	}
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	if m == 0 || n == 0 {
		return 0
	}

	getIdx := func(i, j int) int { return i*n + j }
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	edges := make([][3]int, 0, m*n*2) // (a,b,w)
	for i := range m {
		for j := range n {
			u := getIdx(i, j)
			if j+1 < n {
				v := getIdx(i, j+1)
				edges = append(edges, [3]int{u, v, abs(heights[i][j] - heights[i][j+1])})
			}
			if i+1 < m {
				v := getIdx(i+1, j)
				edges = append(edges, [3]int{u, v, abs(heights[i][j] - heights[i+1][j])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	uf := NewUnionFind(m * n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
		if uf.Find(0) == uf.Find(m*n-1) {
			return edge[2]
		}
	}
	return 0
}
