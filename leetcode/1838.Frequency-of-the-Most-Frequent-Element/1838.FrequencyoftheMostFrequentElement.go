package leetcode

import (
	"slices"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	var ans, sum int
	sort.Ints(nums)
	for l, r := 0, 0; r < len(nums); {
		if k >= (r-l)*nums[r]-sum {
			sum += nums[r]
			r++
			ans = max(ans, r-l)
		} else {
			sum -= nums[l]
			l++
		}
	}
	return ans
}

func maxFrequency2(nums []int, k int) int {
	var ans int
	slices.Sort(nums)
	slices.Reverse(nums)

	for l, r := 0, 0; r < len(nums); {
		if diff := nums[l] - nums[r]; diff <= k {
			k -= diff
			r++
			ans = max(ans, r-l)
		} else {
			l++
			k += (r - l) * (nums[l-1] - nums[l])
		}
	}
	return ans
}
