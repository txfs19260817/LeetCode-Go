package _200_Number_of_Islands

func numIslands(grid [][]byte) int {
	ans, visited := 0, make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				ans++
				dfs(i, j, &visited, grid)
			}
		}
	}
	return ans
}

func dfs(i, j int, visited *[][]bool, grid [][]byte) {
	if grid[i][j] == '1' {
		(*visited)[i][j] = true
	}
	if i+1 < len(grid) && !(*visited)[i+1][j] && grid[i+1][j] == '1' {
		dfs(i+1, j, visited, grid)
	}
	if i-1 >= 0 && !(*visited)[i-1][j] && grid[i-1][j] == '1' {
		dfs(i-1, j, visited, grid)
	}
	if j+1 < len(grid[0]) && !(*visited)[i][j+1] && grid[i][j+1] == '1' {
		dfs(i, j+1, visited, grid)
	}
	if j-1 >= 0 && !(*visited)[i][j-1] && grid[i][j-1] == '1' {
		dfs(i, j-1, visited, grid)
	}
}
