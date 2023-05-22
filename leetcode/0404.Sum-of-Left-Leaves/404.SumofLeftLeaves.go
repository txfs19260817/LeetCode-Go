package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var ans int
	var helper func(*TreeNode, bool)
	helper = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if isLeft {
				ans += node.Val
			}
			return
		}
		if node.Left != nil {
			helper(node.Left, true)
		}
		if node.Right != nil {
			helper(node.Right, false)
		}
	}
	helper(root, false)
	return ans
}
