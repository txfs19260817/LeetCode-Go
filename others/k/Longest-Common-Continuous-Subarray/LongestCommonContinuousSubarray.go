package k

func longestCommonContinuousHistory(history1, history2 []string) (ans []string) {
	dp, maxCnt := make([][]int, len(history1)+1), 0
	for i := range dp {
		dp[i] = make([]int, len(history2)+1)
	}
	for i := 1; i <= len(history1); i++ {
		for j := 1; j <= len(history2); j++ {
			if history1[i-1] == history2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if dp[i][j] > maxCnt {
					maxCnt = dp[i][j]
					ans = history1[i-maxCnt : i]
				}
			}
		}
	}
	return
}
