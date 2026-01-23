package leetcode

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	ans, size := make([]int, n), make([]int, n)
	var dfs func(u, parent, depth int)
	dfs = func(u, parent, depth int) {
		ans[0] += depth // depth 为 0 到 x 的距离
		size[u] = 1 // 初始化当前节点的大小为 1
		for _, v := range g[u] { // 遍历当前节点的所有邻居节点
			if v != parent { // 避免访问父节点
				dfs(v, u, depth+1) // u 是 v 的父节点，深度加 1
				size[u] += size[v]  // 累加 u 的儿子 v 的子树大小
			}
		}
	}
	dfs(0, -1, 0) // 从根节点开始遍历(根节点的父节点为 -1，深度为 0)

	var reroot func(u, parent int)
	reroot = func(u, parent int) {
		for _, v := range g[u] { // 遍历当前节点的所有邻居节点
			if v != parent { // 避免访问父节点
				ans[v] = ans[u] + n - 2*size[v] // 更新 v 的距离
				reroot(v, u) // 递归更新 v 的距离
			}
		}
	}
	reroot(0, -1) // 0 没有父节点

	return ans
}
