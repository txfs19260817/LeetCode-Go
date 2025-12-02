package leetcode

import (
	"math"
	"slices"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[src] = 0

	for range k + 1 { // 最多 k 次中转 => 最多 k+1 段航班 => 迭代 k+1 层
		tmp := slices.Clone(dp)
		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]
			if dp[from] != math.MaxInt {
				tmp[to] = min(tmp[to], dp[from]+price)
			}
		}
		dp = tmp
	}
	if dp[dst] == math.MaxInt {
		return -1
	}
	return dp[dst]
}
