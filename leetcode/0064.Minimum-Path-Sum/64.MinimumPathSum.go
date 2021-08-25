package _064_Minimum_Path_Sum

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
