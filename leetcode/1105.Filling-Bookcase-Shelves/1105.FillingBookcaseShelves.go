package leetcode

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := range books {
		dp[i+1] = math.MaxInt
		leftWidth, maxHeight := shelfWidth, 0
		for j := i; j >= 0; j-- {
			w, h := books[j][0], books[j][1]
			leftWidth -= w
			if leftWidth < 0 {
				break
			}
			maxHeight = max(maxHeight, h)
			dp[i+1] = min(dp[i+1], dp[j]+maxHeight)
		}
	}
	return dp[n]
}
