package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	// l, r := trimBST(root.Left, low, high), trimBST(root.Right, low, high)
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left, root.Right = trimBST(root.Left, low, high), trimBST(root.Right, low, high)
	return root
}
