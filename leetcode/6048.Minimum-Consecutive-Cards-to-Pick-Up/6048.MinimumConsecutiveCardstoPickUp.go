package leetcode

func minimumCardPickup(cards []int) int {
	ans, v2i := 1<<30, map[int][]int{}
	for i, v := range cards {
		v2i[v] = append(v2i[v], i)
	}

	for _, indices := range v2i {
		for i := 0; i < len(indices)-1; i++ {
			ans = min(ans, indices[i+1]-indices[i])
		}
	}

	if ans == 1<<30 {
		return -1
	}
	return ans + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
