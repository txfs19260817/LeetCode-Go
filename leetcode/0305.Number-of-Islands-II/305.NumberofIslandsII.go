package leetcode

const WATER = -1

type UnionFind struct {
	parent, size   []int
	count, maxSize int // Follow-up: calculate the largest island size for each step
}

func NewUnionFind(cnt int) UnionFind {
	parent := make([]int, cnt)
	for i := range parent {
		parent[i] = WATER
	}
	return UnionFind{parent: parent, size: make([]int, cnt)} // Note: count is init with 0
}

func (uf *UnionFind) Union(p, q int) {
	p, q = uf.Find(p), uf.Find(q)
	if p == q {
		return
	}
	// Ensure the size of p is larger
	if uf.size[p] < uf.size[q] {
		p, q = q, p
	}
	uf.parent[q] = p
	uf.size[p] += uf.size[q]
	uf.count--
	uf.maxSize = max(uf.maxSize, uf.size[p])
}

func (uf *UnionFind) Find(p int) int {
	if p != uf.parent[p] {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) IsLand(p int) bool {
	return uf.parent[p] != WATER
}

func (uf *UnionFind) AddLand(p int) {
	if uf.IsLand(p) {
		return
	}
	uf.parent[p] = p
	uf.size[p] = 1
	uf.count++
	uf.maxSize = max(uf.maxSize, 1)
}

func numIslands2(m int, n int, positions [][]int) ([]int, []int) {
	ans, maxSizes := make([]int, len(positions)), make([]int, len(positions))
	uf := NewUnionFind(m * n)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i, p := range positions {
		x, y := p[0], p[1]
		idx := x*n + y
		uf.AddLand(idx)

		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			nIdx := nx*n + ny
			if nx >= 0 && nx < m && ny >= 0 && ny < n && uf.IsLand(nIdx) {
				uf.Union(idx, nIdx)
			}
		}

		ans[i], maxSizes[i] = uf.count, uf.maxSize
	}
	return ans, maxSizes
}
