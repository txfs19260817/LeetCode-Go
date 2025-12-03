package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// averageOfSubtree returns the number of nodes where the value of the node is equal to the average of the values in its subtree.
func averageOfSubtree(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode) (sum int, num int)
	dfs = func(node *TreeNode) (sum int, num int) {
		if node == nil {
			return 0, 0
		}
		if node.Left == nil && node.Right == nil {
			ans++
			return node.Val, 1
		}
		lSum, lNum := dfs(node.Left)
		rSum, rNum := dfs(node.Right)
		sum = lSum + node.Val + rSum
		num = lNum + 1 + rNum
		if sum/num == node.Val {
			ans++
		}
		return sum, num
	}
	dfs(root)
	return ans
}
