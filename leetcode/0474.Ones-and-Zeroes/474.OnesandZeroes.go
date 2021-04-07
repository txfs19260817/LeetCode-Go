package _474_Ones_and_Zeroes

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		var n0, n1 int
		for _, c := range s {
			if c == '0' {
				n0++
			} else {
				n1++
			}
		}
		for i := m; i >= n0; i-- {
			for j := n; j >= n1; j-- {
				if v := 1 + dp[i-n0][j-n1]; dp[i][j] < v {
					dp[i][j] = v
				}
			}
		}
	}
	return dp[m][n]
}
