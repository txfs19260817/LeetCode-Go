package uber

import "math"

// Node defines an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// SumAllValues implements Part 1: Sum of all node values
func SumAllValues(root *Node) int {
	var dfs func(node *Node) int
	dfs = func(node *Node) int {
		if node == nil {
			return 0
		}
		sum := node.Val
		for _, child := range node.Children {
			sum += dfs(child)
		}
		return sum
	}
	return dfs(root)
}

// MaxPathSum implements Part 2: Max path value (Root to Leaf)
func MaxPathSum(root *Node) int {
	sum, _ := MaxPathValues(root)
	return sum
}

// MaxPathValues implements Part 3: Return the node values of the path with max sum
func MaxPathValues(root *Node) (sum int, path []int) {
	if root == nil {
		return
	}
	if len(root.Children) == 0 {
		return root.Val, []int{root.Val}
	}
	bestVal, bestPath := math.MinInt, []int{}
	for _, child := range root.Children {
		childSum, childPath := MaxPathValues(child)
		if childSum > bestVal {
			bestVal, bestPath = childSum, childPath
		}
	}
	return bestVal + root.Val, append([]int{root.Val}, bestPath...)
}
