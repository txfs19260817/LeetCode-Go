package _462_Minimum_Moves_to_Equal_Array_Elements_II

import "sort"

func minMoves2(nums []int) int {
	var ans, mid int
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		mid = (nums[len(nums)/2] + nums[len(nums)/2-1]) / 2
	} else {
		mid = nums[len(nums)/2]
	}
	for _, n := range nums {
		ans += abs(mid - n)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
