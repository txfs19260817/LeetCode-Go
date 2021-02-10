package _107_Binary_Tree_Level_Order_Traversal_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ans = append(ans, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			ans[i] = append(ans[i], q[j].Val)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}
