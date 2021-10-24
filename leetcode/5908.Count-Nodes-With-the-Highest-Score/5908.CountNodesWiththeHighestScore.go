package _908_Count_Nodes_With_the_Highest_Score

func countHighestScoreNodes(parents []int) int {
	g, subTreeSize := make([][]int, len(parents)), make([]int, len(parents))
	for i, parent := range parents {
		if parent != -1 {
			g[parent] = append(g[parent], i)
		}
	}
	// sub-tree size
	var getSize func(node int)
	getSize = func(node int) {
		subTreeSize[node]++
		for _, child := range g[node] {
			getSize(child)
			subTreeSize[node] += subTreeSize[child]
		}
	}
	getSize(0)
	// scores
	var computeScore func(node, remainSize int)
	var ans, maxScore int
	computeScore = func(node, remainSize int) {
		score := remainSize
		for _, child := range g[node] {
			score *= subTreeSize[child]
		}
		if score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		for _, child := range g[node] {
			computeScore(child, len(parents)-subTreeSize[child])
		}
	}
	computeScore(0, 1)
	return ans
}
