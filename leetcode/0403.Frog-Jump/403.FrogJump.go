package _403_Frog_Jump

type pair struct {
	unit, k int
}

func canCross(stones []int) bool {
	if len(stones) < 2 || stones[1] != 1 {
		return false
	}
	// optimization: the frog cannot reach the goal if there exists i that stones[i]-stones[i-1] > i
	for i := 1; i < len(stones); i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	unit2idx := map[int]int{}
	for i, u := range stones {
		unit2idx[u] = i
	}
	visited := map[pair]bool{} // remove duplicates
	stack := []pair{{3, 2}, {2, 1}, {1, 0}}
	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if unit2idx[pos.unit] == len(stones)-1 {
			return true
		}
		if pos.k == 0 || unit2idx[pos.unit] == 0 {
			continue
		}
		for i := -1; i <= 1; i++ {
			p := pair{pos.unit + pos.k + i, pos.k + i}
			if !visited[p] {
				visited[p] = true
				stack = append(stack, p)
			}
		}
	}
	return false
}
