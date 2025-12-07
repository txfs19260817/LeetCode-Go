package leetcode

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	children := make([][]int, nodes)
	for i := 1; i < nodes; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	var dfs func(int) (int, int)
	dfs = func(node int) (sum int, num int) {
		sum, num = value[node], 1
		for _, c := range children[node] {
			cs, cn := dfs(c)
			sum += cs
			num += cn
		}

		if sum == 0 {
			return 0, 0
		}
		return
	}
	_, ans := dfs(0)
	return ans
}
