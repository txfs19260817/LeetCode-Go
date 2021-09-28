package _947_Most_Stones_Removed_with_Same_Row_or_Column

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

func removeStones(stones [][]int) int {
	uf := NewUnionFind(len(stones))
	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				uf.Union(i, j)
			}
		}
	}
	return len(stones) - uf.Count
}
