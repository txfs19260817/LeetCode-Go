package _695_Max_Area_of_Island

func maxAreaOfIsland(grid [][]int) int {
	var ans int
	var dfs func(int, int, *int)
	dfs = func(i int, j int, cnt *int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		*cnt++
		dfs(i+1, j, cnt)
		dfs(i, j+1, cnt)
		dfs(i-1, j, cnt)
		dfs(i, j-1, cnt)
	}
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == 1 {
				var cnt int
				dfs(i, j, &cnt)
				if ans < cnt {
					ans = cnt
				}
			}
		}
	}
	return ans
}
