package _399_Evaluate_Division

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

func (u UnionFind) Same(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ans, op2idx, idx := make([]float64, len(queries)), map[string]int{}, 0
	// build dict (operator: index)
	for _, equation := range equations {
		for _, s := range equation {
			if _, ok := op2idx[s]; !ok {
				op2idx[s] = idx
				idx++
			}
		}
	}
	// build bi-directed graph (adjacency matrix)
	mat := make([][]float64, len(op2idx))
	for i := range mat {
		mat[i] = make([]float64, len(op2idx))
		for j := range mat[i] {
			if i == j {
				mat[i][j] = 1.
			} else {
				mat[i][j] = -1.
			}
		}
	}
	for i, equation := range equations {
		from, to := op2idx[equation[0]], op2idx[equation[1]]
		mat[from][to], mat[to][from] = values[i], 1/values[i]
	}
	// build UF to help check connectivity quickly
	u := NewUnionFind(len(op2idx))
	for _, equation := range equations {
		u.Union(op2idx[equation[0]], op2idx[equation[1]])
	}
	// query
	visited := make([][]bool, len(mat))
	for i := range visited {
		visited[i] = make([]bool, len(mat[0]))
	}
	var dfs func(ans *float64, start, target int, path []float64) bool
	dfs = func(ans *float64, start, target int, path []float64) bool {
		if start == target {
			for _, f := range path {
				*ans *= f
			}
			return true
		}
		for i := 0; i < len(mat[start]); i++ {
			if visited[start][i] || mat[start][i] == -1. {
				continue
			}
			path, visited[start][i] = append(path, mat[start][i]), true
			found := dfs(ans, i, target, path)
			path, visited[start][i] = path[:len(path)-1], false
			if found {
				return true
			}
		}
		return false
	}
	for i, query := range queries {
		from, fromOk := op2idx[query[0]]
		to, toOk := op2idx[query[1]]
		if !(fromOk && toOk && u.Same(from, to)) {
			ans[i] = -1.
			continue
		}
		ans[i] = 1.
		if from == to {
			continue
		}
		dfs(&ans[i], from, to, []float64{})
	}
	return ans
}
