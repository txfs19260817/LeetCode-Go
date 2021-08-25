package _120_Triangle

func minimumTotal(triangle [][]int) int {
	ans, dp := 1<<31, make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = triangle[i][j]
			switch j {
			case 0:
				dp[i][j] += dp[i-1][0]
			case len(dp[i]) - 1:
				dp[i][j] += dp[i-1][len(dp[i-1])-1]
			default:
				dp[i][j] += min(dp[i-1][j-1], dp[i-1][j])
			}
		}
	}
	for _, n := range dp[len(dp)-1] {
		ans = min(ans, n)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
