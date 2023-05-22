package leetcode

func getFood(grid [][]byte) int {
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var q [][2]int
	for i, r := range grid {
		for j, v := range r {
			if v == '#' {
				q = append(q, [2]int{i, j})
			}
		}
	}
	for step := 1; len(q) > 0; step++ {
		length := len(q)
		for i := 0; i < length; i++ {
			r, c := q[0][0], q[0][1]
			q = q[1:]
			for _, d := range dir {
				if x, y := r+d[0], c+d[1]; x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
					if grid[x][y] == '*' {
						return step
					}
					if grid[x][y] == 'O' {
						grid[x][y] = 'X'
						q = append(q, [2]int{x, y})
					}
				}
			}
		}
	}
	return -1
}
