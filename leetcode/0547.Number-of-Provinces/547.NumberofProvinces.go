package leetcode

type UnionFind struct {
	parent, rank []int
	Count        int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{Count: n, parent: make([]int, n), rank: make([]int, n)}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (u *UnionFind) Find(p int) int {
	root := p
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for p != u.parent[p] { // compress
		temp := u.parent[p]
		u.parent[p] = root
		p = temp
	}
	return root
}

func (u *UnionFind) Union(p, q int) {
	pRoot, qRoot := u.Find(p), u.Find(q)
	if pRoot == qRoot {
		return
	}
	if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		u.parent[pRoot] = qRoot
		if u.rank[pRoot] == u.rank[qRoot] {
			u.rank[qRoot]++
		}
	}
	u.Count--
}

func findCircleNum(isConnected [][]int) int {
	uf := NewUnionFind(len(isConnected))
	for i, line := range isConnected {
		for j := i + 1; j < len(line); j++ {
			if line[j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count
}
