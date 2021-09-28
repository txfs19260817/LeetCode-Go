package _684_Redundant_Connection

type UnionFind struct {
	Parent, rank []int
	count        int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{count: n, Parent: make([]int, n), rank: make([]int, n)}
	for i := range u.Parent {
		u.Parent[i] = i
	}
	return u
}

func (u *UnionFind) Find(p int) int {
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

func (u *UnionFind) Union(p, q int) {
	p, q = u.Find(p), u.Find(q)
	if p == q {
		return
	}
	if u.rank[p] < u.rank[q] {
		u.Parent[p] = q
	} else {
		u.Parent[q] = p
		if u.rank[p] == u.rank[q] {
			u.rank[p]++
		}
	}
	u.count--
}

func findRedundantConnection(edges [][]int) []int {
	sameVal := func(arr []int) bool {
		m := map[int]bool{arr[0]: true}
		for _, a := range arr[1:] {
			if !m[a] {
				return false
			}
		}
		return true
	}
	var ans []int
	for i := len(edges) - 1; i >= 0; i-- {
		uf := NewUnionFind(len(edges) + 1) // cuz `n == edges.length`
		for j, edge := range edges {
			if i == j {
				continue
			}
			uf.Union(edge[0], edge[1])
		}
		for i := 1; i <= len(edges); i++ {
			uf.Parent[i] = uf.Find(i)
		}
		if sameVal(uf.Parent[1:]) {
			ans = edges[i]
			break
		}
	}
	return ans
}
