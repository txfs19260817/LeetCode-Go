package leetcode

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // order matters
	visited, dist := make([][]bool, m), make([][]int, m)
	for i := range m {
		visited[i] = make([]bool, n)
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = m * n
		}
	}
	dist[0][0] = 0

	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y := p[0], p[1]
		if visited[x][y] {
			continue
		}
		visited[x][y] = true

		for i, d := range dirs {
			newDist, dirMatch := dist[x][y], i+1 == grid[x][y]
			if !dirMatch {
				newDist++
			}
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && newDist < dist[nx][ny] {
				dist[nx][ny] = newDist
				if dirMatch {
					q = append([][2]int{{nx, ny}}, q...)
				} else {
					q = append(q, [2]int{nx, ny})
				}
			}
		}
	}
	return dist[m-1][n-1]
}
