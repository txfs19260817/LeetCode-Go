package leetcode

func minEdgeReversals(n int, edges [][]int) []int {
	type pair struct{ to, dir int }
	g := make([][]pair, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], pair{v, 1})
		g[v] = append(g[v], pair{u, -1})
	}

	ans := make([]int, n)
	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		for _, v := range g[u] {
			if v.to != parent {
				if v.dir < 0 {
					ans[0]++
				}
				dfs(v.to, u)
			}
		}
	}
	dfs(0, -1)

	var reroot func(u, parent int)
	reroot = func(u, parent int) {
		for _, v := range g[u] {
			if v.to != parent {
				ans[v.to] = ans[u] + v.dir
				reroot(v.to, u)
			}
		}
	}
	reroot(0, -1)

	return ans
}
