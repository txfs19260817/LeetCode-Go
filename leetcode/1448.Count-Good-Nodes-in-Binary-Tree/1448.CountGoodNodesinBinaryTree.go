package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, maxi int)
	dfs = func(node *TreeNode, maxi int) {
		if node == nil {
			return
		}
		if node.Val >= maxi {
			maxi = node.Val
			ans++
		}
		dfs(node.Left, maxi)
		dfs(node.Right, maxi)
	}
	dfs(root, root.Val)
	return ans
}
