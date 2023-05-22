package leetcode

func uniquePathsIII(grid [][]int) int {
	ans, visited := 0, make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	allVisited := func() bool {
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] != -1 && grid[i][j] != 2 && !visited[i][j] {
					return false
				}
			}
		}
		return true
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || visited[i][j] || grid[i][j] == -1 {
			return
		}
		if grid[i][j] == 2 {
			if allVisited() {
				ans++
			}
			return
		}
		visited[i][j] = true
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
		visited[i][j] = false
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	return ans
}
