package _095_Unique_Binary_Search_Trees_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	for i := start; i <= end; i++ {
		ls, rs := helper(start, i-1), helper(i+1, end)
		for _, l := range ls {
			for _, r := range rs {
				allTrees = append(allTrees, &TreeNode{Val: i, Left: l, Right: r})
			}
		}
	}
	return allTrees
}
