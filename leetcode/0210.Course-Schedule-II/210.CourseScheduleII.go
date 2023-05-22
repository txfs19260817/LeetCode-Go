package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	ans, inDegrees, g := make([]int, 0, numCourses), make([]int, numCourses), make([][]int, numCourses)
	for _, pr := range prerequisites {
		g[pr[1]] = append(g[pr[1]], pr[0])
		inDegrees[pr[0]]++
	}
	var q []int
	for i, degree := range inDegrees {
		if degree == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		ans = append(ans, node)
		for _, next := range g[node] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				q = append(q, next)
			}
		}
	}
	if len(ans) != numCourses {
		return nil
	}
	return ans
}
