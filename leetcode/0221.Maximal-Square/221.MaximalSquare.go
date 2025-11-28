package leetcode

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

func maximalSquare2(matrix [][]byte) int {
	n := len(matrix[0])
	ans, dp := 0, [2][]int{make([]int, n+1), make([]int, n+1)}
	for i, row := range matrix {
		for j, v := range row {
			if v == '1' {
				dp[(i+1)%2][j+1] = min(dp[i%2][j], dp[i%2][j+1], dp[(i+1)%2][j]) + 1
				ans = max(ans, dp[(i+1)%2][j+1])
			} else {
				dp[(i+1)%2][j+1] = 0
			}
		}
	}
	return ans * ans
}
