package k

func hasCommonAncestor(edges [][]int, x, y int) bool {
	nodes, child2parent := map[int]bool{}, map[int]map[int]bool{} // child to distinct parents
	for _, pair := range edges {
		parent, child := pair[0], pair[1]
		nodes[parent], nodes[child] = true, true
		if _, ok := child2parent[child]; ok {
			child2parent[child][parent] = true
		} else {
			child2parent[child] = map[int]bool{parent: true}
		}
	}
	xParents, yParents := getAllAncestors(x, child2parent), getAllAncestors(y, child2parent)
	for p := range xParents {
		if yParents[p] {
			return true
		}
	}
	return false
}

func getAllAncestors(x int, child2parent map[int]map[int]bool) map[int]bool {
	ans, childrenQueue := map[int]bool{}, []int{x}
	for len(childrenQueue) > 0 {
		var next []int
		child := childrenQueue[0]
		childrenQueue = childrenQueue[1:]
		parents := child2parent[child]
		for p := range parents {
			ans[p] = true
			next = append(next, p)
		}
		childrenQueue = next
	}
	return ans
}
