package _300_Longest_Increasing_Subsequence

import "sort"

func lengthOfLIS(nums []int) int {
	var seq []int
	for _, n := range nums {
		idx := sort.SearchInts(seq, n)
		if idx == len(seq) {
			seq = append(seq, n)
		} else {
			seq[idx] = n
		}
	}
	return len(seq)
}

func lengthOfLIS1(nums []int) int {
	ans, dp := 1, make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
