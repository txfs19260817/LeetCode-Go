package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	l, r := tree2str(root.Left), tree2str(root.Right)
	ans += "(" + l + ")" // left child cannot be omitted
	if len(r) == 0 {
		return ans
	}
	return ans + "(" + r + ")"
}
