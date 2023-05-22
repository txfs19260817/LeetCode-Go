package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumAverageSubtree(root *TreeNode) float64 {
	var ans float64
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (sum int, num int) {
		if node == nil {
			return 0, 0
		}
		if node.Left == nil && node.Right == nil { // leaf
			if float64(node.Val) > ans { // remember to update the ans
				ans = float64(node.Val)
			}
			return node.Val, 1
		}
		lSum, lNum := dfs(node.Left)
		rSum, rNum := dfs(node.Right)
		sum, num = node.Val+lSum+rSum, 1+lNum+rNum
		if avg := float64(sum) / float64(num); avg > ans {
			ans = avg
		}
		return
	}
	dfs(root)
	return ans
}
