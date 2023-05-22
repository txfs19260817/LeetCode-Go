package leetcode

func calculateMinimumHP(dungeon [][]int) int {
	const MAX = 1 << 30
	dp := make([][]int, len(dungeon)+1)
	for i := range dp {
		dp[i] = make([]int, len(dungeon[0])+1)
		for j := range dp[i] {
			dp[i][j] = MAX
		}
	}
	dp[len(dungeon)-1][len(dungeon[0])], dp[len(dungeon)][len(dungeon[0])-1] = 1, 1
	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			if need := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]; need > 0 {
				dp[i][j] = need
			} else {
				dp[i][j] = 1
			}
		}
	}
	return dp[0][0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
