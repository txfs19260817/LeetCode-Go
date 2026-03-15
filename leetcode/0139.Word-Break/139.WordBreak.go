package leetcode

func wordBreak(s string, wordDict []string) bool {
	m, maxLen := map[string]bool{}, 0
	for _, w := range wordDict {
		m[w] = true
		maxLen = max(maxLen, len(w))
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for _, w := range wordDict {
			if i < len(w) {
				continue
			}
			if !dp[i] {
				dp[i] = dp[i-len(w)] && s[i-len(w):i] == w
				if dp[i] {
					break
				}
			}
		}
	}
	return dp[len(dp)-1]
}
