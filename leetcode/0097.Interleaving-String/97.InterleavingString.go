package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s2)+1)
	dp[0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p] // dp[i][j] = dp[i-1][j] || s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p] // dp[i][j] = dp[i][j-1] || s2[j-1] == s3[p]
			}
		}
	}
	return dp[len(s2)]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
