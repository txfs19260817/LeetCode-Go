package leetcode

// DP
func maxProfit(prices []int) int {
	buy, sell := 0, -1<<31 // hold, release
	for _, p := range prices {
		buy = max(buy, sell+p)
		sell = max(sell, -1*p)
	}
	return buy
}

// find max
func maxProfit1(prices []int) int {
	ans, curMin := 0, prices[0]
	for sellIdx := 1; sellIdx < len(prices); sellIdx++ {
		curMin = min(curMin, prices[sellIdx])
		ans = max(ans, prices[sellIdx]-curMin)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
