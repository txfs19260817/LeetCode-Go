package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	m, parents, children := map[int]*TreeNode{}, map[int]bool{}, map[int]bool{}
	var root *TreeNode
	for _, description := range descriptions {
		iParent, iChild, isLeft := description[0], description[1], description[2] == 1
		p, c := &TreeNode{Val: iParent}, &TreeNode{Val: iChild}
		if n, ok := m[iParent]; ok {
			p = n
		}
		if n, ok := m[iChild]; ok {
			c = n
		}

		if isLeft {
			p.Left = c
		} else {
			p.Right = c
		}

		m[iParent], m[iChild] = p, c
		parents[iParent], children[iChild] = true, true
	}
	for k := range children {
		if parents[k] {
			delete(parents, k)
		}
	}
	for k := range parents {
		root = m[k]
	}
	return root
}
