package leetcode

func maxCoins(nums []int) int {
	nums = append(append([]int{1}, nums...), 1)
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	for l := 2; l < len(nums); l++ {
		for i := 0; i < len(nums)-l; i++ {
			j := i + l
			for k := i + 1; k < j; k++ {
				if v := dp[i][k] + nums[i]*nums[k]*nums[j] + dp[k][j]; v > dp[i][j] {
					dp[i][j] = v
				}
			}
		}
	}
	return dp[0][len(nums)-1]
}
