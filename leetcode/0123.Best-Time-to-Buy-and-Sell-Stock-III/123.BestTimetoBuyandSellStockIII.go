package _122_Best_Time_to_Buy_and_Sell_Stock_II

// DP
func maxProfit(prices []int) int {
	K := 2
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, K)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for k := K; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0], dp[i][k][1] = 0, -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][K-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
