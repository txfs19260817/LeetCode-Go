package _221_Maximal_Square

func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp, prev, maxLen := make([]int, cols+1), 0, 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(dp[j], min(dp[j-1], prev)) + 1
				maxLen = max(maxLen, dp[j])
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}
	return maxLen * maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
