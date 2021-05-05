package _554_Brick_Wall

func leastBricks(wall [][]int) int {
	pos2edges := map[int]int{}
	for _, widths := range wall {
		var i int
		for _, w := range widths[:len(widths)-1] {
			i += w
			pos2edges[i]++
		}
	}
	var maxEdges int
	for _, e := range pos2edges {
		if maxEdges < e {
			maxEdges = e
		}
	}
	return len(wall) - maxEdges
}
