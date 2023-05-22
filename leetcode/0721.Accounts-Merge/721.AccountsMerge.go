package leetcode

import "sort"

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

func (u *UnionFind) GetGroups() [][]int {
	groups, root2idx := make([][]int, 0, u.count), map[int][]int{}
	for i, p := range u.parent {
		root := u.Find(p)
		root2idx[root] = append(root2idx[root], i)
	}
	for _, indices := range root2idx {
		groups = append(groups, indices)
	}
	return groups
}

func accountsMerge(accounts [][]string) [][]string {
	u, name2ids, idx2emailSet := NewUnionFind(len(accounts)), map[string][]int{}, make([]map[string]bool, len(accounts))
	// set methods
	hasInter := func(a, b map[string]bool) bool {
		for k := range a {
			if b[k] {
				return true
			}
		}
		return false
	}
	mergeSet := func(a, b map[string]bool) map[string]bool {
		for k := range a {
			b[k] = true
		}
		return b
	}
	set2Slice := func(a map[string]bool) []string {
		s := make([]string, 0, len(a))
		for k := range a {
			s = append(s, k)
		}
		return s
	}
	// obtain: 1. name to indices; 2. index to email (set)
	for i, account := range accounts {
		name2ids[account[0]] = append(name2ids[account[0]], i)
		idx2emailSet[i] = map[string]bool{}
		for _, email := range account[1:] {
			idx2emailSet[i][email] = true
		}
	}
	// union indices who share any same account
	for _, indices := range name2ids {
		for i := 0; i < len(indices); i++ {
			for j := i + 1; j < len(indices); j++ {
				if idxI, idxJ := indices[i], indices[j]; hasInter(idx2emailSet[idxI], idx2emailSet[idxJ]) {
					u.Union(idxI, idxJ)
				}
			}
		}
	}
	ans := make([][]string, 0, u.count)
	// for each group of indices
	for _, indices := range u.GetGroups() {
		// merge emails of each index & get distinct email slice
		emailSet := map[string]bool{}
		for _, idx := range indices {
			emailSet = mergeSet(emailSet, idx2emailSet[idx])
		}
		emails := set2Slice(emailSet)
		sort.Strings(emails)            // the question required
		name := accounts[indices[0]][0] // indices should share the same name
		ans = append(ans, append([]string{name}, emails...))
	}

	return ans
}
