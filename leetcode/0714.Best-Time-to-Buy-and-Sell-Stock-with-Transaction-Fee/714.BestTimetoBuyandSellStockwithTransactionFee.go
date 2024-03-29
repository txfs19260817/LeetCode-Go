package leetcode

func maxProfit(prices []int, fee int) int {
	buy, sell := 0, -1<<31 // hold, release
	for _, p := range prices {
		buy_1 := buy
		buy = max(buy, sell+p)
		sell = max(sell, buy_1-fee-p)
	}
	return buy
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
