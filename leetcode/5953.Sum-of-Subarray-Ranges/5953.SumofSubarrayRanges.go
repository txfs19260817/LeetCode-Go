package leetcode

func subArrayRanges(nums []int) int64 {
	var ans int64
	for i := 0; i < len(nums); i++ {
		maxV, minV := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			maxV, minV = max(maxV, nums[j]), min(minV, nums[j])
			ans += int64(maxV) - int64(minV)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
