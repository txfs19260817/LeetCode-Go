package _538_Convert_BST_to_Greater_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	ans, sum := root, 0
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += root.Val
		root.Val = sum
		root = root.Left
	}
	return ans
}
