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
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				p = append(p, q[i].Left)
			}
			if q[i].Right != nil {
				p = append(p, q[i].Right)
			}
		}
		if len(q) > 1 && len(p) > 0 {
			foundX, foundY := find(p, x), find(p, y)
			if foundX && !foundY || !foundX && foundY {
				return false
			}
			if foundX && foundY {
				for _, node := range q {
					if node.Left == nil || node.Right == nil {
						continue
					}
					if node.Left.Val == x && node.Right.Val == y || node.Left.Val == y && node.Right.Val == x {
						return false
					}
				}
				return true
			}
		}
		q = p
	}
	return false
}

func find(p []*TreeNode, target int) bool {
	for _, node := range p {
		if node.Val == target {
			return true
		}
	}
	return false
}
