package leetcode

func findShortestCycle(n int, edges [][]int) int {
	type pair struct{ curNode, fromNode int }
	ans := n + 1
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	dist := make([]int, n) // min path: start -> i
	for start := 0; start < n; start++ {
		for i := range dist {
			dist[i] = -1
		}
		dist[start] = 0

		q := []pair{{start, -1}}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range g[u.curNode] {
				if dist[v] < 0 {
					dist[v] = dist[u.curNode] + 1
					q = append(q, pair{v, u.curNode})
				} else if v != u.fromNode {
					ans = min(ans, dist[u.curNode]+dist[v]+1)
				}
			}
		}
	}

	if ans == n+1 {
		return -1
	}
	return ans
}
