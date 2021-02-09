package _103_Binary_Tree_Zigzag_Level_Order_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			res[i] = append(res[i], q[j].Val)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			for l, r := 0, len(res[i])-1; l < r; l, r = l+1, r-1 {
				res[i][l], res[i][r] = res[i][r], res[i][l]
			}
		}
	}
	return res
}
