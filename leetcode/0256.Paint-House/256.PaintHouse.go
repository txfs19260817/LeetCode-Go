package _256_Paint_House

func minCost(costs [][]int) int {
	pre := costs[len(costs)-1]
	for i := len(costs) - 2; i >= 0; i-- {
		cur := []int{costs[i][0], costs[i][1], costs[i][2]}
		cur[0] += min(pre[1], pre[2])
		cur[1] += min(pre[0], pre[2])
		cur[2] += min(pre[0], pre[1])
		pre = cur
	}
	return min(min(pre[0], pre[1]), pre[2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
