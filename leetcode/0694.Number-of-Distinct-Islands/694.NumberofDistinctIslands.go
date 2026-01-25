package leetcode

import "strings"

func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var curIsland strings.Builder

	var dfs func(r, c int, step byte)
	dfs = func(r, c int, step byte) {
		// Fix bounds check: 'n' was compared with 'm' (n < 0 || n >= m) which is incorrect and weird.
		// It should be c < 0 || c >= n
		if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || grid[r][c] == 0 {
			return
		}
		visited[r][c] = true
		curIsland.WriteByte(step)
		
		// Use distinct characters for directions to uniquely identify path
		dfs(r-1, c, 'U')
		dfs(r+1, c, 'D')
		dfs(r, c+1, 'R')
		dfs(r, c-1, 'L')
		
		// Backtracking step to uniquely identify structure (important!)
		curIsland.WriteByte('B') 
	}

	islandSet := map[string]bool{}
	for i, row := range grid {
		for j, v := range row {
			// Fix check: v == '1' assumes ascii, but grid is int.
			if v == 1 && !visited[i][j] {
				curIsland.Reset()
				dfs(i, j, 'S') // Start
				islandSet[curIsland.String()] = true
			}
		}
	}
	return len(islandSet)
}
