package _319_Number_of_Operations_to_Make_Network_Connected

type UnionFind struct {
	parent, rank []int
	count        int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{count: n, parent: make([]int, n), rank: make([]int, n)}
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
		t := u.parent[p]
		u.parent[p] = root
		p = t
	}
	return root
}

func (u *UnionFind) Union(p, q int) (alreadyUnionized bool) {
	p, q = u.Find(p), u.Find(q)
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
	u.count--
	return false
}

func makeConnected(n int, connections [][]int) int {
	u := NewUnionFind(n)
	var cable int
	for _, conn := range connections {
		if u.Union(conn[0], conn[1]) {
			cable++
		}
	}
	if cable >= u.count-1 {
		return u.count - 1
	}
	return -1
}
