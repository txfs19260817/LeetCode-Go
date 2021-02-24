package _673_Number_of_Longest_Increasing_Subsequence

func findNumberOfLIS(nums []int) int {
	ans, dp, count := 0, make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i], count[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
	}
	longest := 0
	for _, d := range dp {
		if d > longest {
			longest = d
		}
	}
	for i, c := range count {
		if dp[i] == longest {
			ans += c
		}
	}
	return ans
}
