package _317_Shortest_Distance_from_All_Buildings

func shortestDistance(grid [][]int) int {
	ans, buildings, dirs := 1<<30, 0, [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	distances := make([][][]int, len(grid))
	for i := range distances {
		distances[i] = make([][]int, len(grid[0]))
		for j := range distances[i] {
			distances[i][j] = make([]int, 2) // (distance, count of buildings)
		}
	}

	bfs := func(startX, startY int) {
		q, visited := [][]int{{startX, startY}}, make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		for level := 0; len(q) > 0; level++ {
			p := make([][]int, 0, len(q)*4)
			for i := range q {
				r, c := q[i][0], q[i][1]
				if grid[r][c] == 0 {
					distances[r][c][0] += level
					distances[r][c][1]++
				}

				for _, d := range dirs {
					if R, C := r+d[0], c+d[1]; R >= 0 && R < len(grid) && C >= 0 && C < len(grid[0]) && !visited[R][C] && grid[R][C] == 0 {
						visited[R][C] = true
						p = append(p, []int{R, C})
					}
				}
			}
			q = p
		}
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				buildings++
				bfs(i, j)
			}
		}
	}

	for _, row := range distances {
		for _, d := range row {
			if d[1] == buildings && ans > d[0] {
				ans = d[0]
			}
		}
	}

	if ans == 1<<30 { // not found
		ans = -1
	}
	return ans
}

// TLE
func shortestDistance1(grid [][]int) int {
	ans, buildings, dirs := 1<<30, 0, [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				buildings++
			}
		}
	}

	bfs := func(startX, startY int) (dist int) {
		q, visited, met := [][]int{{startX, startY}}, make([][]bool, len(grid)), 0
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		for level := 0; len(q) > 0; level++ {
			p := make([][]int, 0, len(q)*4)
			for i := range q {
				r, c := q[i][0], q[i][1]
				if grid[r][c] == 1 {
					dist += level
					met++
					if met == buildings { // return dist if we found all targets
						return
					}
					continue
				}

				for _, d := range dirs {
					if R, C := r+d[0], c+d[1]; R >= 0 && R < len(grid) && C >= 0 && C < len(grid[0]) && !visited[R][C] && grid[R][C] != 2 {
						visited[R][C] = true
						p = append(p, []int{R, C})
					}
				}
			}
			q = p
		}
		return -1 // met != buildings
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				if dist := bfs(i, j); dist != -1 && dist < ans {
					ans = dist
				}
			}
		}
	}

	if ans == 1<<30 { // not found
		ans = -1
	}
	return ans
}
