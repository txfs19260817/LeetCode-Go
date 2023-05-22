package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	a := root
	for {
		if p.Val < a.Val && q.Val < a.Val {
			a = a.Left
		} else if p.Val > a.Val && q.Val > a.Val {
			a = a.Right
		} else {
			break
		}
	}
	return a
}
