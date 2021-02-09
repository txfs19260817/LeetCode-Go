package _098_Validate_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := -1 << 63
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if inorder >= root.Val {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
