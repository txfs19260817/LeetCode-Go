package leetcode

func minimumTime(n int, relations [][]int, time []int) int {
	g := make([][]int, n)
	inDeg := make([]int, n)
	for _, relation := range relations {
		prevCourse, nextCourse := relation[0]-1, relation[1]-1
		g[prevCourse] = append(g[prevCourse], nextCourse)
		inDeg[nextCourse]++
	}

	var q []int
	for c, d := range inDeg {
		if d == 0 {
			q = append(q, c)
		}
	}

	var ans int
	dp := make([]int, n)
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		dp[c] += time[c]
		ans = max(ans, dp[c])

		for _, nextCourse := range g[c] {
			dp[nextCourse] = max(dp[nextCourse], dp[c])

			if inDeg[nextCourse]--; inDeg[nextCourse] == 0 {
				q = append(q, nextCourse)
			}
		}
	}
	return ans
}
