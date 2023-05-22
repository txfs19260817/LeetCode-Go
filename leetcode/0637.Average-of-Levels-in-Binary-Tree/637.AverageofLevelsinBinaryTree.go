package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var levels [][]int
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		levels = append(levels, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			levels[i] = append(levels[i], q[j].Val)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	ans := make([]float64, len(levels))
	for i, level := range levels {
		for _, num := range level {
			ans[i] += float64(num)
		}
		ans[i] /= float64(len(level))
	}
	return ans
}
