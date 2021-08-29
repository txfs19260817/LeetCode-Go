package _062_Unique_Paths

import "math/big"

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[n-1]
}

func uniquePaths2(m, n int) int { // c_(m+n-2)^(n-1) or c_(m+n-2)^(m-1)
	return int(new(big.Int).Binomial(int64(m+n-2), int64(m-1)).Int64())
}
