package leetcode

// DP
func maxProfit(prices []int) int {
	buy, sell := 0, -1<<31 // hold, release
	for _, p := range prices {
		buy = max(buy, sell+p)
		sell = max(sell, buy-p)
		//temp := buy
		//buy = max(buy, sell+p)
		//sell = max(sell, temp-p)
	}
	return buy
}

// Greedy
func maxProfit1(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
