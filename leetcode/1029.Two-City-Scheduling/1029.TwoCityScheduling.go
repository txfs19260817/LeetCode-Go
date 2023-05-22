package leetcode

import "sort"

func twoCitySchedCost(costs [][]int) int {
	// sort by the company's gain by sending a person to A than B
	sort.Slice(costs, func(i, j int) bool { return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1] })
	ans, n := 0, len(costs)/2
	for i := 0; i < n; i++ {
		ans += costs[i][0] + costs[i+n][1] // first half to A; second half to B
	}
	return ans
}
