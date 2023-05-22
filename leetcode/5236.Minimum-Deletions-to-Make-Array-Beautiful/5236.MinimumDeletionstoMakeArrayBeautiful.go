package leetcode

func minDeletion(nums []int) int {
	var stack []int
	for i := 0; i < len(nums); i += 2 {
		if len(stack)%2 == 1 {
			if nums[i] == stack[len(stack)-1] {
				i--
				continue
			}
			stack = append(stack, nums[i])
			if i+1 < len(nums) {
				stack = append(stack, nums[i+1])
			}
		} else {
			stack = append(stack, nums[i])
			if i+1 < len(nums) && nums[i] != nums[i+1] {
				stack = append(stack, nums[i+1])
			}
		}
	}
	return len(nums) - len(stack)/2*2
}
