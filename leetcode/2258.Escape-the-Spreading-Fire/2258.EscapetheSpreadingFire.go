package leetcode

func maximumMinutes(grid [][]int) int {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	bfs := func(q [][2]int) (int, int, int) { // 返回三个数，分别表示到达安全屋/安全屋左边/安全屋上边的最短时间
		times := make([][]int, m)
		for i := range times {
			times[i] = make([]int, n)
			for j := range n {
				times[i][j] = -1 // -1 表示未访问
			}
		}
		for _, x := range q {
			times[x[0]][x[1]] = 0
		}

		for l := 1; len(q) > 0; l++ { // 每次循环向外扩展一圈
			var p [][2]int
			for _, x := range q {
				for _, d := range dirs { // 枚举上下左右四个方向
					ni, nj := x[0]+d[0], x[1]+d[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 0 && times[ni][nj] == -1 {
						times[ni][nj] = l
						p = append(p, [2]int{ni, nj})
					}
				}
			}
			q = p
		}
		return times[m-1][n-1], times[m-2][n-1], times[m-1][n-2]
	}

	manTime, m1, m2 := bfs([][2]int{{0, 0}})
	if manTime < 0 {
		return -1
	}

	var firePos [][2]int
	for i, r := range grid {
		for j, v := range r {
			if v == 1 {
				firePos = append(firePos, [2]int{i, j})
			}
		}
	}
	fireTime, f1, f2 := bfs(firePos)
	if fireTime < 0 {
		return 1e9
	}

	d := fireTime - manTime
	if d < 0 {
		return -1
	}

	if m1 != -1 && f1-m1 > d || // 安全屋左边相邻格子，人比火先到
		m2 != -1 && f2-m2 > d { // 安全屋上边相邻格子，人比火先到
		return d
	}
	return d - 1
}
