package _515_Find_Largest_Value_in_Each_Tree_Row

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var p []*TreeNode
		maxV := -1 << 31
		for j := 0; j < len(q); j++ {
			if q[j].Val > maxV {
				maxV = q[j].Val
			}
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		ans = append(ans, maxV)
		q = p
	}
	return ans
}
