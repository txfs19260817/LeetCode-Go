package leetcode

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		next := make([]int, 0, len(nums)/2)
		for i := 1; i < len(nums); i += 2 {
			if len(next)%2 == 0 {
				next = append(next, min(nums[i-1], nums[i]))
			} else {
				next = append(next, max(nums[i-1], nums[i]))
			}
		}
		nums = next
	}
	return nums[0]
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
