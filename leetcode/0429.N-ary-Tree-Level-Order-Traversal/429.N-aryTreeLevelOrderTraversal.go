package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	for q := []*Node{root}; len(q) > 0; {
		var p []*Node
		var level []int
		for _, node := range q {
			level = append(level, node.Val)
			for _, child := range node.Children {
				p = append(p, child)
			}
		}
		result = append(result, level)
		q = p
	}
	return result
}
