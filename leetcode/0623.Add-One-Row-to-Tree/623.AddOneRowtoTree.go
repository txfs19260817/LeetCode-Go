package _623_Add_One_Row_to_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	q := []*TreeNode{root}
	for i := 0; i < depth-2 && len(q) > 0; i++ {
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	for j := 0; j < len(q); j++ {
		q[j].Left = &TreeNode{val, q[j].Left, nil}
		q[j].Right = &TreeNode{val, nil, q[j].Right}
	}
	return root
}
