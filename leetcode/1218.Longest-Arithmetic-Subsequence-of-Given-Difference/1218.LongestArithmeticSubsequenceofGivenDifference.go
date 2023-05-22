package leetcode

func longestSubsequence(arr []int, difference int) int {
	ans, dp := 1, map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans
}
