package _000_Minimum_Cost_to_Merge_Stones

// https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/solution/yi-dong-you-yi-dao-nan-yi-bu-bu-shuo-ming-si-lu-he/
func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	prefix := make([]int, n+1)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + stones[i-1]
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for l := k; l <= n; l++ {
		for i := 1; i+l-1 <= n; i++ {
			j := i + l - 1
			dp[i][j] = 1 << 30
			for p := i; p < j; p += k - 1 {
				dp[i][j] = min(dp[i][j], dp[i][p]+dp[p+1][j])
			}
			if (j-i)%(k-1) == 0 {
				dp[i][j] += prefix[j] - prefix[i-1]
			}
		}
	}
	return dp[1][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
