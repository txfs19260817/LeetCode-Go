package leetcode

import "sort"

func sortArrayByParity(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return !(nums[i]%2 == 1 && nums[j]%2 == 0) })
	return nums
}

/*
If it is (0, 1), then everything is correct: i++ and j--.
If it is (1, 0), we swap them, so they are correct, then continue.
If it is (0, 0), only the i place is correct, so we i++ and continue.
If it is (1, 1), only the j place is correct, so we j-- and continue.
*/
func sortArrayByParity2(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]%2 > nums[j]%2 { // 1, 0
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]%2 == 0 {
			i++
		}
		if nums[j]%2 == 1 {
			j--
		}
	}
	return nums
}
