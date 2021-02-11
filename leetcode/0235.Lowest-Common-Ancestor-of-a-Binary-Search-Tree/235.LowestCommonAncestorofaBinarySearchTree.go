package _235_Lowest_Common_Ancestor_of_a_Binary_Search_Tree

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
