package leetcode

type UnionFind struct {
	parent, rank []int
	Count        int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{Count: n, parent: make([]int, n), rank: make([]int, n)}
	for i := range u.parent {
		u.parent[i] = i
	}
	return u
}

func (u *UnionFind) Find(p int) int {
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
	u.Count--
}

func findRedundantConnection(edges [][]int) []int {
	var ans []int
	for i := len(edges) - 1; i >= 0; i-- {
		uf := NewUnionFind(len(edges) + 1) // cuz `n == edges.length` + unused [0]
		for j, edge := range edges {
			if i == j {
				continue
			}
			uf.Union(edge[0], edge[1])
		}
		if uf.Count == 2 { // 0 + other groups
			ans = edges[i]
			break
		}
	}
	return ans
}
