package _152_Maximum_Product_Subarray

func maxProduct(nums []int) int {
	ans, maxF, minF := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		ma, mi := maxF, minF
		maxF = max(nums[i], max(ma*nums[i], mi*nums[i]))
		minF = min(nums[i], min(ma*nums[i], mi*nums[i]))
		ans = max(ans, maxF)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
