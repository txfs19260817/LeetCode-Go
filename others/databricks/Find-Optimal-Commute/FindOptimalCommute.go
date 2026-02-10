package databricks

// FindOptimalCommute returns the name of the fastest transportation mode
// from S to D in the grid. Ties on time are broken by least cost.
//
// Grid cells: "S" = source, "D" = destination, "X" = roadblock,
// "1" = Walk, "2" = Bike, "3" = Car, "4" = Train.
// costs[i] and times[i] correspond to mode i+1 (0-indexed).
func FindOptimalCommute(grid [][]string, costs []int, times []int) string {
	rows := len(grid)
	if rows == 0 {
		return ""
	}
	cols := len(grid[0])

	// Locate source and destination.
	var sr, sc, dr, dc int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case "S":
				sr, sc = r, c
			case "D":
				dr, dc = r, c
			}
		}
	}

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// bfs returns the shortest path length (in blocks) from S to D
	// using only cells that match modeStr, or -1 if unreachable.
	bfs := func(modeStr string) int {
		visited := make([][]bool, rows)
		for i := range visited {
			visited[i] = make([]bool, cols)
		}
		type cell struct{ r, c int }
		queue := []cell{{sr, sc}}
		visited[sr][sc] = true

		for steps := 0; len(queue) > 0; steps++ {
			var next []cell
			for _, p := range queue {
				if p.r == dr && p.c == dc {
					return steps
				}
				for _, d := range dirs {
					nr, nc := p.r+d[0], p.c+d[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || visited[nr][nc] {
						continue
					}
					if grid[nr][nc] == modeStr || grid[nr][nc] == "D" {
						visited[nr][nc] = true
						next = append(next, cell{nr, nc})
					}
				}
			}
			queue = next
		}
		return -1
	}

	modeStrs := [4]string{"1", "2", "3", "4"}
	modeNames := [4]string{"Walk", "Bike", "Car", "Train"}

	bestMode := ""
	bestTime, bestCost := 1<<62, 1<<62

	for i, ms := range modeStrs {
		dist := bfs(ms)
		if dist < 0 {
			continue
		}
		t := dist * times[i]
		c := dist * costs[i]
		if t < bestTime || (t == bestTime && c < bestCost) {
			bestTime = t
			bestCost = c
			bestMode = modeNames[i]
		}
	}

	return bestMode
}
