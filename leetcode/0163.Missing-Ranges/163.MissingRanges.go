package leetcode

import "strconv"

func findMissingRanges(nums []int, lower int, upper int) []string {
	var ans []string
	if len(nums) == 0 {
		if lower == upper {
			ans = append(ans, strconv.Itoa(lower))
		} else {
			ans = append(ans, strconv.Itoa(lower)+"->"+strconv.Itoa(upper))
		}
		return ans
	}
	if lower != nums[0] {
		nums = append([]int{lower - 1}, nums...) // lower - 1: inclusive boundary, offset the +1 below
	}
	if upper != nums[len(nums)-1] {
		nums = append(nums, upper+1) // upper+1: inclusive boundary, offset the -1 below
	}
	for i := 0; i < len(nums)-1; i++ {
		if newLower, newUpper := nums[i]+1, nums[i+1]-1; newUpper == newLower {
			ans = append(ans, strconv.Itoa(newLower))
		} else if newLower < newUpper {
			ans = append(ans, strconv.Itoa(newLower)+"->"+strconv.Itoa(newUpper))
		}
	}
	return ans
}
