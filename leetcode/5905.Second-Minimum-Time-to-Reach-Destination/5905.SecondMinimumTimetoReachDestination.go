package _905_Second_Minimum_Time_to_Reach_Destination

func secondMinimum(n int, edges [][]int, time int, change int) int {
	ans, g := 0, make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	q, shortestDist, reachable := []int{0}, make([]int, n), make([]bool, n)
	for i := 1; i < len(shortestDist); i++ { // the shortest distance from start to vertex i
		shortestDist[i] = 1 << 31
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			if shortestDist[u]+1 < shortestDist[v] {
				shortestDist[v] = shortestDist[u] + 1
				q = append(q, v)
			} else if shortestDist[u] == shortestDist[v] || (reachable[u] && shortestDist[u] <= shortestDist[v]) {
				if !reachable[v] {
					q = append(q, v)
					reachable[v] = true
				}
			}
		}
	}
	dist := shortestDist[n-1] + 2
	if reachable[n-1] {
		dist--
	}
	for dist > 0 {
		dist--
		ans += time
		if dist > 0 && ans%(2*change) >= change {
			ans = (ans/(2*change) + 1) * 2 * change
		}
	}
	return ans
}
