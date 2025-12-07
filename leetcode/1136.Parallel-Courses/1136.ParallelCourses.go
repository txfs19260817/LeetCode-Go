package leetcode

func minimumSemesters(n int, relations [][]int) int {
	var ans int
	g, inDeg := make([][]int, n), make([]int, n)
	for _, relation := range relations {
		pre, next := relation[0]-1, relation[1]-1
		g[pre] = append(g[pre], next)
		inDeg[next]++
	}

	var q []int
	for c, d := range inDeg {
		if d == 0 {
			q = append(q, c)
		}
	}
	learnt := len(q)

	for len(q) > 0 {
		ans++
		var p []int
		for _, c := range q {
			for _, nc := range g[c] {
				if inDeg[nc]--; inDeg[nc] == 0 {
					learnt++
					p = append(p, nc)
				}
			}
		}
		q = p
	}
	if learnt != n {
		return -1
	}
	return ans
}
