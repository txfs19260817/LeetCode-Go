package _312_Minimum_Insertion_Steps_to_Make_a_String_Palindrome

func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for l := 2; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			j := i + l - 1
			dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
