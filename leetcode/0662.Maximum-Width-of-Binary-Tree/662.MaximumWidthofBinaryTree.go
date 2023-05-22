package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	type item struct {
		idx  int
		node *TreeNode
	}
	ans, q := 1, []item{{0, root}}
	for len(q) > 0 {
		if length := q[len(q)-1].idx - q[0].idx + 1; length > ans {
			ans = length
		}
		var p []item
		for i := 0; i < len(q); i++ {
			if q[i].node.Left != nil {
				p = append(p, item{idx: 2 * q[i].idx, node: q[i].node.Left})
			}
			if q[i].node.Right != nil {
				p = append(p, item{idx: 2*q[i].idx + 1, node: q[i].node.Right})
			}
		}
		q = p
	}
	return ans
}
