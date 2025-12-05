package leetcode

import "math"

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1 // 注意这里加一了
		for j := n; j > 0; j-- {
			dp[j] = min(dp[j], dp[max(j-t, 0)]+c)
		}
	}
	return dp[n]
}
