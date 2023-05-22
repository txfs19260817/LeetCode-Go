package leetcode

func maxProfit(prices []int) int {
	buy, sell, cooldown := 0, -1<<31, 0 // hold, release, dp[i-2]
	for _, p := range prices {
		temp := buy
		buy = max(buy, sell+p)
		sell = max(sell, cooldown-p)
		cooldown = temp
	}
	return buy
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
