package _993_Cousins_in_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	q := []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		var ParentX, ParentY *TreeNode
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				if q[i].Left.Val == x {
					ParentX = q[i]
				} else if q[i].Left.Val == y {
					ParentY = q[i]
				}
				p = append(p, q[i].Left)
			}
			if q[i].Right != nil {
				if q[i].Right.Val == x {
					ParentX = q[i]
				} else if q[i].Right.Val == y {
					ParentY = q[i]
				}
				p = append(p, q[i].Right)
			}
		}
		if ParentX != nil && ParentY != nil {
			return ParentX != ParentY
		}
		if ParentX != nil || ParentY != nil {
			break // return false
		}
		q = p
	}
	return false
}
