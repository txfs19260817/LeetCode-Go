package leetcode

func minStartValue(nums []int) int {
	minValue := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		if minValue > nums[i] {
			minValue = nums[i]
		}
	}
	if minValue >= 0 {
		return 1
	}
	return 1 - minValue
}
