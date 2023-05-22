package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	for p := root; p != nil; p = p.Right {
		if p.Left == nil {
			continue
		}
		temp := p.Right
		p.Right = p.Left
		p.Left = nil
		q := p
		for q.Right != nil {
			q = q.Right
		}
		q.Right = temp
	}
}
