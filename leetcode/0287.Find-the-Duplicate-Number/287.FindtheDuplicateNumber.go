package _287_Find_the_Duplicate_Number

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
			continue
		}
		if nums[nums[i]-1] == nums[i] {
			return nums[i]
		}
		nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
	}
	return nums[len(nums)-1]
}

func findDuplicate1(nums []int) int {
	s, f := 0, 0
	for s, f = nums[s], nums[nums[f]]; s != f; s, f = nums[s], nums[nums[f]] {
	}
	s = 0
	for ; s != f; s, f = nums[s], nums[f] {
	}
	return s
}
