package leetcode

func numEnclaves(grid [][]int) int {
	var ans int
	for i := range grid {
		dfs(grid, i, 0)
		dfs(grid, i, len(grid[0])-1)
	}
	for i := range grid[0] {
		dfs(grid, 0, i)
		dfs(grid, len(grid)-1, i)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}
