package _220_Contains_Duplicate_III

import "sort"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	copyNums := append([]int{}, nums...)
	sortInterval := k + 1
	if k+1 > len(nums) {
		sortInterval = len(nums)
	}
	sort.Ints(nums[:sortInterval])
	for i := 1; i <= k && i < len(nums); i++ {
		if abs(int64(nums[i])-int64(nums[i-1])) <= int64(t) {
			return true
		}
	}
	nums = copyNums
	for i := k + 1; i < len(nums); i++ {
		for j := i - 1; j > i-k-1; j-- {
			if abs(int64(nums[i])-int64(nums[j])) <= int64(t) {
				return true
			}
		}
	}
	return false
}

func abs(a int64) int64 {
	if a < 0 {
		return -1 * a
	}
	return a
}
