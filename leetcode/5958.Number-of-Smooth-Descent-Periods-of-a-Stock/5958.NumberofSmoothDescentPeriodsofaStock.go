package leetcode

func getDescentPeriods(prices []int) int64 {
	ans, dp := int64(0), make([]int, len(prices))
	dp[0] = 1
	for i := 1; i < len(prices); i++ {
		if prices[i]+1 == prices[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	for _, d := range dp {
		ans += int64(d)
	}
	return ans
}
