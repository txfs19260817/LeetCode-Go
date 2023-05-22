package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var levels [][]int
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		levels = append(levels, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			levels[i] = append(levels[i], q[j].Val)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	for _, level := range levels {
		res = append(res, level[len(level)-1])
	}
	return res
}
