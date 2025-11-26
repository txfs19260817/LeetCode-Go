package leetcode

import "math"

func minXor(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	xor := 0
	for i := 1; i <= n; i++ {
		xor ^= nums[i-1]
		dp[i] = xor
	}

	for p := 2; p <= k; p++ {
		for i := n; i >= p; i-- {
			dp[i] = math.MaxInt
			xor := 0
			for j := i - 1; j >= p-1; j-- {
				xor ^= nums[j]
				dp[i] = min(dp[i], max(dp[j], xor))
			}
		}
	}
	return dp[n]
}
