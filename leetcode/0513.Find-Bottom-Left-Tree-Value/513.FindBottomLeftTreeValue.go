package _513_Find_Bottom_Left_Tree_Value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var ans int
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		if len(p) == 0 {
			ans = q[0].Val
			break
		}
		q = p
	}
	return ans
}
