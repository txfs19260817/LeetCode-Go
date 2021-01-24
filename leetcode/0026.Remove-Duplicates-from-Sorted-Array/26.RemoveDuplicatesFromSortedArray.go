package _026_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	s, f := 0, 1
	for s < len(nums)-1 {
		for nums[s] == nums[f] {
			f++
			if f == len(nums) {
				return s + 1
			}
		}
		s++
		nums[s] = nums[f]
	}
	return s + 1
}
