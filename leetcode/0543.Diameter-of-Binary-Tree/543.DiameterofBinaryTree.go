package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		l, r := dfs(node.Left)+1, dfs(node.Right)+1
		ans = max(ans, l+r) // update max LENGTH = max left depth + max right depth
		return max(l, r)    // return max DEPTH
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
