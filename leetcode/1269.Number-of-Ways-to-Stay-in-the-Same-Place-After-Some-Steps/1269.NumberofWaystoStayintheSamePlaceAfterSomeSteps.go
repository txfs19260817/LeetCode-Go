package _269_Number_of_Ways_to_Stay_in_the_Same_Place_After_Some_Steps

func numWays(steps int, arrLen int) int {
	const MOD = 1e9 + 7
	minCols := arrLen
	// If a pointer from index 0 wants to go back with `steps` steps,
	// it can only reach as far as index steps/2.
	if c := steps/2 + 1; c < minCols {
		minCols = c
	}
	dp := make([]int, minCols)
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		dpNext := make([]int, minCols)
		for j := 0; j < minCols; j++ {
			dpNext[j] = dp[j]
			if j-1 >= 0 {
				dpNext[j] = (dpNext[j] + dp[j-1]) % MOD
			}
			if j+1 < minCols {
				dpNext[j] = (dpNext[j] + dp[j+1]) % MOD
			}
		}
		dp = dpNext
	}
	return dp[0]
}
