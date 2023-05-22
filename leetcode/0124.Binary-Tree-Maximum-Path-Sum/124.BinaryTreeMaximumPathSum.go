package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

const intMin = -1 << 31

func maxPathSum(root *TreeNode) int {
	ans = intMin
	helper(root)
	return ans
}

func helper(root *TreeNode) int {
	if root == nil {
		return intMin
	}
	cand1, cand2, tri := intMin, intMin, intMin
	l, r := helper(root.Left), helper(root.Right)
	if root.Left != nil && root.Right != nil {
		tri = root.Val + l + r
	}
	if root.Left != nil {
		cand1 = l + root.Val
	}
	if root.Right != nil {
		cand2 = r + root.Val
	}
	ans = max(ans, max(tri, max(root.Val, max(cand1, cand2))))
	return max(root.Val, max(cand1, cand2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
