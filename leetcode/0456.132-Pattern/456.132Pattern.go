package _456_132_Pattern

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	mins, stack := make([]int, len(nums)), make([]int, 0, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		mins[i] = min(mins[i-1], nums[i])
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if mins[i] >= nums[i] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] <= mins[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			return true
		}
		stack = append(stack, nums[i])
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
