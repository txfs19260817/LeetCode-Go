package leetcode

func orangesRotting(grid [][]int) int {
	var ans int
	var q [][2]int
	for i, line := range grid { // 0: black, 2: gray, 1: white
		for j, orange := range line {
			if orange == 2 {
				q = append(q, [2]int{i, j})
			}
		}
	}
	for len(q) > 0 {
		var nextQ [][2]int
		for _, point := range q {
			i, j := point[0], point[1]
			grid[i][j] = 0
			if i-1 >= 0 && j >= 0 && i-1 < len(grid) && j < len(grid[0]) && grid[i-1][j] == 1 {
				nextQ = append(nextQ, [2]int{i - 1, j})
				grid[i-1][j] = 2
			}
			if i+1 >= 0 && j >= 0 && i+1 < len(grid) && j < len(grid[0]) && grid[i+1][j] == 1 {
				nextQ = append(nextQ, [2]int{i + 1, j})
				grid[i+1][j] = 2
			}
			if i >= 0 && j-1 >= 0 && i < len(grid) && j-1 < len(grid[0]) && grid[i][j-1] == 1 {
				nextQ = append(nextQ, [2]int{i, j - 1})
				grid[i][j-1] = 2
			}
			if i >= 0 && j+1 >= 0 && i < len(grid) && j+1 < len(grid[0]) && grid[i][j+1] == 1 {
				nextQ = append(nextQ, [2]int{i, j + 1})
				grid[i][j+1] = 2
			}
		}
		q = nextQ
		if len(q) > 0 {
			ans++
		}
	}
	for _, line := range grid {
		for _, orange := range line {
			if orange == 1 {
				return -1
			}
		}
	}
	return ans
}
