package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	helper(root, &sum)
	return sum
}

func helper(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	l, r := helper(root.Left, sum), helper(root.Right, sum)
	*sum += abs(l - r)
	return root.Val + l + r
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
