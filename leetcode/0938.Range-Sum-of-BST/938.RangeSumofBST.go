package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var ans int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val >= low && root.Val <= high {
			ans += root.Val
		}
		root = root.Right
	}
	return ans
}
