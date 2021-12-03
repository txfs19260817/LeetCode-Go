package k

func findNodesWithZeroOrOneParent(edges [][]int) (p0, p1 []int) {
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
	for node := range nodes {
		if p := len(child2parent[node]); p == 0 {
			p0 = append(p0, node)
		} else if p == 1 {
			p1 = append(p1, node)
		}
	}
	return
}
