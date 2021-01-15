package _324_Wiggle_Sort_II

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)
	res := make([]int, len(nums))

	mid := len(nums) / 2
	if len(nums)%2 == 1 {
		mid++
	}

	for i, j := mid-1, 0; i >= 0; i, j = i-1, j+2 {
		res[j] = nums[i]
	}
	for i, j := len(nums)-1, 1; i >= mid; i, j = i-1, j+2 {
		res[j] = nums[i]
	}

	copy(nums, res)
}
