package _872_Leaf_Similar_Trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves1, leaves2 []int
	var getLeaves func(*TreeNode, *[]int)
	getLeaves = func(root *TreeNode, leaves *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*leaves = append(*leaves, root.Val)
		}
		getLeaves(root.Left, leaves)
		getLeaves(root.Right, leaves)
	}
	getLeaves(root1, &leaves1)
	getLeaves(root2, &leaves2)
	if len(leaves1) == len(leaves2) {
		for i, n := range leaves1 {
			if leaves2[i] != n {
				return false
			}
		}
		return true
	}
	return false
}
