package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[r+1] == nums[r] {
					r--
				}
				for l < r && nums[l-1] == nums[l] {
					l++
				}
			}
		}
	}
	return res
}
