package _712_Minimum_ASCII_Delete_Sum_for_Two_Strings

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := len(s1) - 1; i >= 0; i-- {
		dp[i][len(s2)] = dp[i+1][len(s2)] + int(s1[i])
	}
	for j := len(s2) - 1; j >= 0; j-- {
		dp[len(s1)][j] = dp[len(s1)][j+1] + int(s2[j])
	}
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}
	return dp[0][0]
}

func minimumDeleteSum1(s1 string, s2 string) int {
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var recursion func(i, j int) int
	recursion = func(i, j int) int {
		var ans int
		if i == len(s1) {
			for ; j < len(s2); j++ {
				ans += int(s2[j])
			}
			return ans
		}
		if j == len(s2) {
			for ; i < len(s1); i++ {
				ans += int(s1[i])
			}
			return ans
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		if s1[i] == s2[j] {
			dp[i][j] = recursion(i+1, j+1)
		} else {
			dp[i][j] = min(int(s1[i])+recursion(i+1, j), int(s2[j])+recursion(i, j+1))
		}
		return dp[i][j]
	}
	return recursion(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
