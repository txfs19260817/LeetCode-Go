package _053_Maximum_Subarray

func maxSubArray(nums []int) int {
	ans, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = nums[i] + max(cur, 0)
		ans = max(ans, cur)
	}
	return ans
}

func maxSubArray1(nums []int) int {
	ans, dp := nums[0], make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	for _, d := range dp {
		ans = max(ans, d)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
