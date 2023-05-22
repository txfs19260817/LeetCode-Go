package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var lh, rh int
	for node := root; node.Left != nil; node = node.Left {
		lh++
	}
	for node := root; node.Right != nil; node = node.Right {
		rh++
	}
	if lh == rh {
		return 1<<(lh+1) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
