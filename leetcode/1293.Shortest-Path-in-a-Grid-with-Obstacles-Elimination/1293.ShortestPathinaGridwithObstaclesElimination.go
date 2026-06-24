package leetcode

type state struct {
	i, j, remainingK, step int
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	q := []state{{0, 0, k, 0}}
	seen := map[state]bool{{0, 0, k, 0}: true}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.i == m-1 && curr.j == n-1 {
			return curr.step
		}
		for _, dir := range dirs {
			nextI, nextJ := curr.i+dir[0], curr.j+dir[1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n {
				nextK := curr.remainingK - grid[nextI][nextJ]
				nextState := state{nextI, nextJ, nextK, curr.step + 1}
				if nextStateWithoutStep := (state{nextI, nextJ, nextK, 0}); nextK >= 0 && !seen[nextStateWithoutStep] {
					seen[nextStateWithoutStep] = true
					q = append(q, nextState)
				}
			}
		}
	}
	return -1
}
