package _375_Guess_Number_Higher_or_Lower_II

import "math"

func getMoneyAmount(n int) int {
	cache := map[[2]int]int{}
	dp := make([]int, max(3, n)+1)
	dp[2], dp[3] = 1, 2
	for N := 4; N <= n; N++ {
		dp[N] = 1 << 31
		for i := 2; i <= N-1; i++ {
			if isBalance(i, 1, N) {
				dp[N] = min(dp[N], i+max(dp[i-1], recursion(i+1, N, cache)))
			}
		}
	}
	return dp[n]
}

func recursion(low, high int, cache map[[2]int]int) int {
	switch high - low {
	case 0:
		return 0
	case 1:
		return low
	case 2:
		return low + 1
	default:
		if v, ok := cache[[2]int{low, high}]; ok {
			return v
		}
		ans := 1 << 31
		for i := low + 1; i <= high-1; i++ {
			if isBalance(i, low, high) {
				ans = min(ans, i+max(recursion(low, i-1, cache), recursion(i+1, high, cache)))
			}
		}
		cache[[2]int{low, high}] = ans
		return ans
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalance(i, low, high int) bool {
	return math.Abs(math.Min(math.Ceil(math.Log2(float64(i-low))), 1)-math.Min(math.Ceil(math.Log2(float64(high-i))), 1)) < 2
}

func getMoneyAmount2(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for j := 2; j <= n; j++ {
		for i := j - 1; i > 0; i-- {
			if i+1 == j {
				dp[i][j] = i
				continue
			}
			dp[i][j] = 1 << 31
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}
