package _413_Arithmetic_Slices

func numberOfArithmeticSlices(nums []int) int {
	var ans, cur int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cur++
			ans += cur
		} else {
			cur = 0
		}
	}
	return ans
}

// DP version of the solution above
func numberOfArithmeticSlices2(nums []int) int {
	ans, dp := 0, make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i-2] == nums[i]-nums[i-1] {
			dp[i] = 1 + dp[i-1]
			ans += dp[i]
		}
	}
	return ans
}
