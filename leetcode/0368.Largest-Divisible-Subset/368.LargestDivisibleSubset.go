package leetcode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	dp := make([]int, len(nums))  // dp[i] points to last index in ans
	cnt := make([]int, len(nums)) // ans's length
	maxCnt, maxCntIdx := 1, 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		dp[i], cnt[i] = i, 1
		for j := i - 1; j >= 0; j-- {
			if (nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0) && cnt[i] < cnt[j]+1 {
				dp[i] = j
				cnt[i] = cnt[j] + 1
			}
		}
		if cnt[i] > maxCnt {
			maxCnt, maxCntIdx = cnt[i], i
		}
	}
	// trace reversely to collect all elements
	ans := make([]int, maxCnt)
	for i := maxCnt - 1; i >= 0; i-- {
		ans[i] = nums[maxCntIdx]
		maxCntIdx = dp[maxCntIdx] // move to previous one
		if maxCntIdx == dp[maxCntIdx] {
			break
		}
	}
	ans[0] = nums[maxCntIdx]
	return ans
}
