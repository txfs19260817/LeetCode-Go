package _576_Out_of_Boundary_Paths

const mod int = 1e9 + 7

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	ans, dp := 0, mat(m, n)
	dp[startRow][startColumn] = 1
	for i := 0; i < maxMove; i++ {
		dpNew := mat(m, n)
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				count := dp[j][k]
				for _, dir := range dirs {
					if j1, k1 := j+dir.x, k+dir.y; j1 >= 0 && j1 < m && k1 >= 0 && k1 < n {
						dpNew[j1][k1] = (dpNew[j1][k1] + count) % mod
					} else {
						ans = (ans + count) % mod
					}
				}
			}
		}
		dp = dpNew
	}
	return ans
}

func mat(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return dp
}
