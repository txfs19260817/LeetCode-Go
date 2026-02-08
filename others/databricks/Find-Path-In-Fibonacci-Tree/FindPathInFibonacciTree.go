package databricks

import "strings"

// fibNodes returns the number of nodes in a Fibonacci tree of the given order.
// fib_nodes(0) = 1, fib_nodes(1) = 1, fib_nodes(n) = 1 + fib_nodes(n-2) + fib_nodes(n-1).
func fibNodes(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1 // fib_nodes(0), fib_nodes(1)
	for i := 2; i <= n; i++ {
		a, b = b, 1+a+b
	}
	return b
}

// pathFromRoot returns the sequence of 'L'/'R' moves from the root of a
// Fibonacci tree (of the given order, with root labeled rootLabel) to target.
func pathFromRoot(order, rootLabel, target int) string {
	var path []byte
	for order > 1 && rootLabel != target {
		leftSize := fibNodes(order - 2)
		rightStart := rootLabel + 1 + leftSize
		if target < rightStart {
			path = append(path, 'L')
			rootLabel = rootLabel + 1
			order = order - 2
		} else {
			path = append(path, 'R')
			rootLabel = rightStart
			order = order - 1
		}
	}
	return string(path)
}

// FindPath returns the path from source to dest in a Fibonacci tree of the
// given order, expressed as a string of 'L', 'R', and 'U' moves.
func FindPath(order, source, dest int) string {
	if source == dest {
		return ""
	}
	pathToSource := pathFromRoot(order, 0, source)
	pathToDest := pathFromRoot(order, 0, dest)

	// Longest common prefix = LCA depth
	common := 0
	for common < len(pathToSource) && common < len(pathToDest) &&
		pathToSource[common] == pathToDest[common] {
		common++
	}

	ups := strings.Repeat("U", len(pathToSource)-common)
	return ups + pathToDest[common:]
}
