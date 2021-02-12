package _337_House_Robber_III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var helper func(*TreeNode) (int, int) // ->cur, below
	helper = func(node *TreeNode) (selected int, notSelected int) {
		if node == nil {
			return 0, 0
		}
		ls, lns := helper(node.Left)
		rs, rns := helper(node.Right)
		selected = node.Val + lns + rns
		notSelected = max(ls, lns) + max(rs, rns)
		return
	}
	return max(helper(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
