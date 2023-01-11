package leetcode

import "sort"

// prefix sum
func minSubArrayLen(s int, nums []int) int {
	ans, preSum := 1<<31, make([]int, len(nums)+1) // prefix sum array has an extra 1st element which is 0
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	for i, p := range preSum {
		target := s + p
		j := sort.SearchInts(preSum, target)
		if j == len(preSum) {
			break
		}
		if preSum[j] >= target {
			ans = min(ans, j-i) // instead of r-l+1
		}
	}

	if ans == 1<<31 {
		return 0
	}
	return ans
}

// sliding window
func minSubArrayLen2(s int, nums []int) int {
	ans, sum := 1<<31, 0
	for l, r := 0, 0; l < len(nums); {
		if sum < s && r < len(nums) {
			sum += nums[r]
			r++
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			ans = min(ans, r-l)
		}
	}
	if ans == 1<<31 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
