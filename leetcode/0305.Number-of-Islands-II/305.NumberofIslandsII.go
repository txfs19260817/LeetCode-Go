package leetcode

type UnionFind struct {
	parent, rank []int
	count        int
}

func NewUnionFind(cnt int) UnionFind {
	parent := make([]int, cnt)
	for i := range parent {
		parent[i] = i
	}
	return UnionFind{parent: parent, rank: make([]int, cnt), count: 0} // Note: count is init with 0
}

func (u *UnionFind) Find(p int) int {
	if p != u.parent[p] {
		u.parent[p] = u.Find(u.parent[p])
	}
	return u.parent[p]
}

func (u *UnionFind) Union(p, q int) {
	p, q = u.Find(p), u.Find(q)
	if p == q {
		return
	}
	if u.rank[p] < u.rank[q] {
		u.parent[p] = q
	} else {
		u.parent[q] = p
		if u.rank[p] == u.rank[q] {
			u.rank[p]++
		}
	}
	u.count--
}

func numIslands2(m int, n int, positions [][]int) []int {
	ans := make([]int, 0, len(positions))
	visited := make(map[int]bool, m*n)
	uf := NewUnionFind(m * n)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, p := range positions {
		x, y := p[0], p[1]
		idx := x*n + y
		if !visited[idx] {
			visited[idx] = true
			uf.count++
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				nIdx := nx*n + ny
				if nx >= 0 && nx < m && ny >= 0 && ny < n && visited[nIdx] {
					uf.Union(idx, nIdx)
				}
			}
		}
		ans = append(ans, uf.count)
	}
	return ans
}
