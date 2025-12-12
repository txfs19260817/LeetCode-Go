package leetcode

import "sort"

func maxValue(events [][]int, k int) int {
	n := len(events)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	for i, e := range events {
		startDay, value := e[0], e[2]
		p := sort.Search(i, func(j int) bool { return events[j][1] >= startDay })
		for j := 1; j <= k; j++ {
			dp[i+1][j] = max(dp[i][j], dp[p][j-1]+value)
		}
	}
	return dp[n][k]
}
