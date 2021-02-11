package _110_Balanced_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh, rh := height(root.Left), height(root.Right)
	if lh == -1 || rh == -1 || abs(lh-rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
