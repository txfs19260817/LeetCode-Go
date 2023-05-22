package leetcode

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	u := &unionFind{make([]int, n), make([]int, n)}
	for i := range u.parent {
		u.parent[i] = i
	}
	return u
}

func (u *unionFind) find(p int) int {
	root := p
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for p != u.parent[p] {
		tmp := u.parent[p]
		u.parent[p] = root
		p = tmp
	}
	return root
}

func (u *unionFind) Union(p, q int) (alreadyUnionized bool) {
	p, q = u.find(p), u.find(q)
	if p == q {
		return true
	}
	if u.rank[p] > u.rank[q] {
		u.parent[q] = p
	} else {
		u.parent[p] = q
		if u.rank[p] == u.rank[q] {
			u.rank[q]++
		}
	}
	return false
}

func containsCycle(grid [][]byte) bool {
	uf := newUnionFind(len(grid) * len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			if j >= 1 && grid[i][j] == grid[i][j-1] && uf.Union(len(grid[0])*i+j, len(grid[0])*i+j-1) { // left
				return true
			}
			if i >= 1 && grid[i][j] == grid[i-1][j] && uf.Union(len(grid[0])*i+j, len(grid[0])*(i-1)+j) { // up
				return true
			}
		}
	}
	return false
}
