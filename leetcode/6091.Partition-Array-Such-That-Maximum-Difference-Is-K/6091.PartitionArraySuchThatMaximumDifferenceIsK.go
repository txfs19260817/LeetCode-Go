package _091_Partition_Array_Such_That_Maximum_Difference_Is_K

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
