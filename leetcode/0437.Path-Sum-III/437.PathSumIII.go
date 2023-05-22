package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	var ans int
	var path []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		var tempSum int
		for i := len(path) - 1; i >= 0; i-- {
			tempSum += path[i]
			if tempSum == sum {
				ans++
			}
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
