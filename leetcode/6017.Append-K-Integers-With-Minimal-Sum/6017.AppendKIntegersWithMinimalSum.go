package leetcode

import "sort"

func minimalKSum(nums []int, k int) int64 {
	ans := int64(0)
	nums = append(nums, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1 && k > 0; i++ {
		start := nums[i] + 1
		n := nums[i+1] - 1 - start + 1
		if n <= 0 {
			continue
		}
		n = min(k, n)
		k -= n
		ans += int64(n)*int64(start) + int64(n)*int64(n-1)/2

	}

	if k > 0 {
		ans += int64(k)*int64(nums[len(nums)-1]+1) + int64(k)*int64(k-1)/2
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
