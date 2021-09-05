package _092_Shortest_Common_Supersequence

func shortestCommonSupersequence(str1 string, str2 string) string {
	ans, common := make([]byte, 0, len(str1)+len(str2)), longestCommonSubsequence(str1, str2)
	var i, j int
	for _, b := range common {
		for ; i < len(str1) && str1[i] != b; i++ {
			ans = append(ans, str1[i])
		}
		for ; j < len(str2) && str2[j] != b; j++ {
			ans = append(ans, str2[j])
		}
		ans = append(ans, b)
		i++
		j++
	}
	if i < len(str1) {
		ans = append(ans, str1[i:]...)
	}
	if j < len(str2) {
		ans = append(ans, str2[j:]...)
	}
	return string(ans)
}

func longestCommonSubsequence(text1 string, text2 string) []byte {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	common := make([]byte, dp[len(text1)][len(text2)])
	for i, j, k := len(text1), len(text2), len(common)-1; dp[i][j] > 0; {
		if text1[i-1] == text2[j-1] {
			common[k] = text1[i-1]
			i, j, k = i-1, j-1, k-1
			continue
		}
		if dp[i-1][j] == dp[i][j] {
			i--
		} else {
			j--
		}
	}
	return common
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
