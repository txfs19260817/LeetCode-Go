package _388_Pizza_With_3n_Slices

func maxSizeSlices(slices []int) int {
	return max(calc(slices[:len(slices)-1]), calc(slices[1:]))
}

func calc(slices []int) int {
	dp := make([][]int, len(slices)+1) // pick j from i nums
	for i := range dp {
		dp[i] = make([]int, (len(slices)+1)/3+1) // n-1+1
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			pick := slices[i-1]
			if i-2 >= 0 {
				pick += dp[i-2][j-1]
			}
			dp[i][j] = max(dp[i-1][j], pick)
		}
	}
	return dp[len(slices)][(len(slices)+1)/3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
