package leetcode

// DFS/Tri-Color
// Time Complexity: O(n+m) for edges + vertices
// Space Complexity: O(n) for colors
func eventualSafeNodes(graph [][]int) []int {
	colors := make([]uint8, len(graph)) // 0: not visited, 1: not sure, 2: confirmed
	var dfs func(int) bool
	dfs = func(i int) bool {
		if colors[i] > 0 {
			return colors[i] == 2
		}
		colors[i] = 1
		for _, to := range graph[i] { // all outgoing edges can reach any of terminal node finally
			if !dfs(to) {
				return false
			}
		}
		colors[i] = 2
		return true
	}
	var ans []int
	for i := range graph {
		if dfs(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

// Topological sorting
// Time Complexity: O(n+m) for edges + vertices
// Space Complexity: O(n+m) for inverse graph
func eventualSafeNodes2(graph [][]int) []int {
	revGraph, inDeg := make([][]int, len(graph)), make([]int, len(graph))
	for i, to := range graph {
		for _, t := range to {
			revGraph[t] = append(revGraph[t], i)
		}
		inDeg[i] = len(to)
	}
	var q []int // push all nodes whose inDeg = 0 to a queue
	for i, ds := range inDeg {
		if ds == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		node := q[0] // remove an inDeg = 0 node from rev graph
		q = q[1:]
		for _, toRev := range revGraph[node] {
			inDeg[toRev]--
			if inDeg[toRev] == 0 { // continue push new inDeg = 0 nodes to queue
				q = append(q, toRev)
			}
		}
	}
	var ans []int // return all inDeg = 0 nodes
	for i, ds := range inDeg {
		if ds == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
