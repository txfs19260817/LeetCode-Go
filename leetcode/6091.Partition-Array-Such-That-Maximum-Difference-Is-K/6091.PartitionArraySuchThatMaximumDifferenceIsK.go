package leetcode

import "sort"

func partitionArray(nums []int, k int) int {
	var ans int
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		i = sort.SearchInts(nums, nums[i]+k+1)
		ans++
	}
	return ans
}
