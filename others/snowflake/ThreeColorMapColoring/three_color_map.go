package snowflake

const (
	Unknown = iota
	Red
	Blue
	Green
)

func threeColor(locations []string, adjacencies [][]string) bool {
	g := map[string][]string{}
	for _, adjacency := range adjacencies {
		u, v := adjacency[0], adjacency[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	colorMap := map[string]int{}
	for _, location := range locations {
		colorMap[location] = Unknown
	}

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(locations) {
			return true
		}

		currentNode := locations[idx]

		for _, c := range []int{Red, Blue, Green} {
			// Check if color c is valid
			isValid := true
			for _, neighbor := range g[currentNode] {
				if colorMap[neighbor] == c {
					isValid = false
					break
				}
			}

			if isValid {
				colorMap[currentNode] = c
				if dfs(idx + 1) {
					return true
				}
				colorMap[currentNode] = Unknown // Backtrack
			}
		}
		return false
	}

	return dfs(0)
}
