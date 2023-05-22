package leetcode

func maximumScore(nums []int, multipliers []int) int {
	ans, dp := -1<<31, make([][]int, len(multipliers)+1)
	for i := range dp {
		dp[i] = make([]int, len(multipliers)+1)
	}
	for k := 1; k <= len(multipliers); k++ { // pick k numbers from `nums`
		for i := 0; i <= k; i++ {
			j := k - i
			if i == 0 {
				dp[i][j] = dp[i][j-1] + multipliers[k-1]*nums[len(nums)-j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + multipliers[k-1]*nums[i-1]
			} else {
				dp[i][j] = max(dp[i-1][j]+multipliers[k-1]*nums[i-1], dp[i][j-1]+multipliers[k-1]*nums[len(nums)-j])
			}
			if k == len(multipliers) {
				ans = max(ans, dp[i][j])
			}
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
