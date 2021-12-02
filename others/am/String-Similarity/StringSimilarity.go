package am

func similarities(s string) []int {
	dp := make([]int, len(s))
	for i, l, r := 1, 0, 0; i < len(s); i++ {
		if i <= r {
			dp[i] = min(dp[i-l], r-i+1)
		}
		for dp[i]+i < len(s) && s[dp[i]] == s[dp[i]+i] {
			dp[i]++
		}
		if dp[i]+i-1 > r {
			l, r = i, dp[i]+i-1
		}
	}
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
