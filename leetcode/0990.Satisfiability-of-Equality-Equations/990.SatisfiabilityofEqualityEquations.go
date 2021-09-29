package _990_Satisfiability_of_Equality_Equations

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

func (u *UnionFind) Union(p, q int) {
	p, q = u.Find(p), u.Find(q)
	if p == q {
		return
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
}

func (u *UnionFind) Same(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func equationsPossible(equations []string) bool {
	ch2idx, idx := map[byte]int{}, 0
	for _, eq := range equations {
		if _, ok := ch2idx[eq[0]]; !ok {
			ch2idx[eq[0]] = idx
			idx++
		}
		if _, ok := ch2idx[eq[3]]; !ok {
			ch2idx[eq[3]] = idx
			idx++
		}
	}
	u := NewUnionFind(idx + 1)
	for _, eq := range equations {
		if eq[1] == '=' {
			u.Union(ch2idx[eq[0]], ch2idx[eq[3]])
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' && u.Same(ch2idx[eq[0]], ch2idx[eq[3]]) {
			return false
		}
	}
	return true
}
