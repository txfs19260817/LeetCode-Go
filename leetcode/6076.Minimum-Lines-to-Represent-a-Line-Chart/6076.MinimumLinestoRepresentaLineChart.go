package leetcode

import "sort"

func minimumLines(stockPrices [][]int) int {
	var ans int
	sort.Slice(stockPrices, func(i, j int) bool { return stockPrices[i][0] < stockPrices[j][0] })
	for i, preDY, preDX := 1, 1, 0; i < len(stockPrices); i++ {
		dy, dx := stockPrices[i][1]-stockPrices[i-1][1], stockPrices[i][0]-stockPrices[i-1][0]
		if dy*preDX != preDY*dx {
			ans++
			preDY, preDX = dy, dx
		}
	}
	return ans
}
