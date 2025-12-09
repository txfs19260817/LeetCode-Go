package leetcode

func minTotalDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var rows, cols []int
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				rows = append(rows, i)
			}
		}
	}
	for j := range n {
		for i := range m {
			if grid[i][j] == 1 {
				cols = append(cols, j)
			}
		}
	}

	mr, mc := rows[len(rows)/2], cols[len(cols)/2]
	return dist(rows, mr) + dist(cols, mc)
}

func dist(points []int, median int) (d int) {
	for _, point := range points {
		d += abs(point - median)
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
