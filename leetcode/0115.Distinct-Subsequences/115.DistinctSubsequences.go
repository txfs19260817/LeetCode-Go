package _115_Distinct_Subsequences

func numDistinct(s string, t string) int {
	ans, dp := 0, make([]int, len(s)) // dp[j] = the num of s[:j] occurs sub t[:i]
	for i := 0; i < len(t); i++ {
		var pre int
		if i == 0 {
			pre = 1 // empty string is a substring
		}
		for j := 0; j < len(s); j++ {
			pre += dp[j] // accumulate substrings
			if s[j] == t[i] {
				dp[j] = pre - dp[j] // t[:i] in s[:j] excluding t[:i] in s[:j-1]
			} else {
				dp[j] = 0 // the intro of s[j] has no contribute to any new sub
			}
		}
	}
	for _, n := range dp {
		ans += n
	}
	return ans
}
