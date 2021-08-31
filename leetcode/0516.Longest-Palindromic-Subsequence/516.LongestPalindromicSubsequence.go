package _516_Longest_Palindromic_Subsequence

func longestPalindromeSubseq(s string) int {
	t, dp := reverse(s), make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(s)][len(s)]
}

func reverse(s string) string {
	c := []byte(s)
	for l, r := 0, len(c)-1; l < r; l, r = l+1, r-1 {
		c[l], c[r] = c[r], c[l]
	}
	return string(c)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeSubseq2(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			} else if i+1 == j && s[i] == s[j] {
				dp[i][j] = 2
			} else {
				if s[i] == s[j] {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
				} else {
					dp[i][j] = max(dp[i][j], max(dp[i+1][j], dp[i][j-1]))
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
