package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, pr := range prerequisites {
		g[pr[0]] = append(g[pr[0]], pr[1])
	}
	bfs := func(node int) bool {
		visited := make([]bool, numCourses)
		visited[node] = true
		for q := g[node]; len(q) > 0; {
			var p []int
			for _, n := range q {
				if n == node {
					return false
				}
				if visited[n] {
					continue
				}
				visited[n] = true
				p = append(p, g[n]...)
			}
			q = p
		}
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !bfs(i) { // starting from `i`, check if it will return to `i`.
			return false
		}
	}
	return true
}
