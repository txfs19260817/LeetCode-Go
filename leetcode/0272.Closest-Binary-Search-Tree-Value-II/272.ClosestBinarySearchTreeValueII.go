package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	var ans []int
	var s []*TreeNode
	for len(s) > 0 || root != nil {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]

		if len(ans) < k {
			ans = append(ans, root.Val)
		} else if math.Abs(float64(root.Val)-target) < math.Abs(float64(ans[0])-target) {
			ans = append(ans[1:], root.Val)
		}

		root = root.Right
	}
	return ans
}
