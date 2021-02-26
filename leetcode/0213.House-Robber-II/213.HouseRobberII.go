package _213_House_Robber_II

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robSingle(nums[:len(nums)-1]), robSingle(nums[1:]))
}

// 198
func robSingle(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
