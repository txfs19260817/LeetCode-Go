package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodes(root *TreeNode) []int {
	var ans []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || node.Left == nil && node.Right == nil {
			return
		}
		if node.Left == nil {
			ans = append(ans, node.Right.Val)
		}
		if node.Right == nil {
			ans = append(ans, node.Left.Val)
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
