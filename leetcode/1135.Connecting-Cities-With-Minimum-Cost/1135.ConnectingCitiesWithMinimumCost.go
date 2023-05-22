package leetcode

import "sort"

type unionFind struct {
	Parent, rank []int
	count        int
}

func newUnionFind(n int) *unionFind {
	u := &unionFind{make([]int, n), make([]int, n), n}
	for i := range u.Parent {
		u.Parent[i] = i
	}
	return u
}

func (u *unionFind) Find(p int) int {
	root := p
	for root != u.Parent[root] {
		root = u.Parent[root]
	}
	for p != u.Parent[p] {
		tmp := u.Parent[p]
		u.Parent[p] = root
		p = tmp
	}
	return root
}
func (u *unionFind) Union(p, q int) (alreadyUnionized bool) {
	p, q = u.Find(p), u.Find(q)
	if p == q {
		return true
	}
	if u.rank[p] > u.rank[q] {
		u.Parent[q] = p
	} else {
		u.Parent[p] = q
		if u.rank[p] == u.rank[q] {
			u.rank[q]++
		}
	}
	u.count--
	return false
}

func minimumCost(n int, connections [][]int) int {
	ans, uf := 0, newUnionFind(n+1) // 1 to n
	sort.Slice(connections, func(i, j int) bool { return connections[i][2] < connections[j][2] })
	for _, c := range connections {
		if !uf.Union(c[0], c[1]) {
			ans += c[2]
		}
	}
	root := uf.Find(1)
	for i := 2; i <= n; i++ {
		if uf.Find(i) != root {
			return -1
		}
	}
	return ans
}
