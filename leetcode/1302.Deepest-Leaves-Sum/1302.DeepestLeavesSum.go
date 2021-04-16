package _302_Deepest_Leaves_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	ans, maxHeight := 0, -1
	var dfs func(node *TreeNode, h int)
	dfs = func(node *TreeNode, h int) {
		if node == nil {
			return
		}
		if h > maxHeight {
			maxHeight = h
			ans = node.Val // update max height along with the deepest leaf
		} else if h == maxHeight {
			ans += node.Val
		}
		dfs(node.Left, h+1)
		dfs(node.Right, h+1)
	}
	dfs(root, 0)
	return ans
}
