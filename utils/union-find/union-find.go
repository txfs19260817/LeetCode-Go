package union_find

// 路径压缩 + 秩优化
type UnionFind struct {
	parent, rank []int
	count        int
}

// Init
func (uf *UnionFind) Init(n int) {
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	uf.count = n
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

// Find
func (uf *UnionFind) Find(p int) int {
	root := uf.parent[p]
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// path compression
	for p != uf.parent[p] {
		temp := uf.parent[p] // Get parent node
		uf.parent[p] = root  // set the current node's parent to root
		p = temp             // handle the next node
	}
	return root
}

// Union
func (uf *UnionFind) Union(p, q int) {
	pRoot, qRoot := uf.Find(p), uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[qRoot] = pRoot
		if uf.rank[pRoot] == uf.rank[qRoot] {
			uf.rank[pRoot]++
		}
	}
	uf.count--
}

// TotalCount
func (uf *UnionFind) TotalCount() int {
	return uf.count
}
