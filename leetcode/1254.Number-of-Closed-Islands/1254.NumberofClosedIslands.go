package leetcode

// similar: 1254, 1020, 130
func closedIsland(grid [][]int) int {
	var ans int
	for i := range grid {
		dfs(grid, i, 0)
		dfs(grid, i, len(grid[0])-1)
	}
	for i := 1; i < len(grid[0])-1; i++ {
		dfs(grid, 0, i)
		dfs(grid, len(grid)-1, i)
	}
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == 0 {
				dfs(grid, i, j)
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
