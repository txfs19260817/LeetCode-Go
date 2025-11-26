package leetcode

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	dp := make([]float64, n+1)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += nums[i-1]
		dp[i] = float64(sum) / float64(i)
	}

	for p := 2; p <= k; p++ {
		for i := n; i >= p; i-- {
			s := 0
			for j := i - 1; j >= p-1; j-- {
				s += nums[j]
				avg := float64(s) / float64(i-j)
				dp[i] = max(dp[i], dp[j]+avg)
			}
		}
	}
	return dp[n]
}
