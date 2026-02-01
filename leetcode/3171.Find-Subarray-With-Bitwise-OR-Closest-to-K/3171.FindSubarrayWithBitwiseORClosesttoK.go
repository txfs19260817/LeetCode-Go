package leetcode

import "math"

func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		ans = min(ans, abs(x-k)) // single number
		for j := i - 1; j >= 0 && (nums[j]|x) != nums[j]; j-- {
			nums[j] |= x
			ans = min(ans, abs(k-nums[j]))
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
