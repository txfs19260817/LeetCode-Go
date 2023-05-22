package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	inRootIdx := indexOf(inorder, postorder[len(postorder)-1])
	inLeft, inRight := inorder[:inRootIdx], inorder[inRootIdx+1:]
	postLeft, postRight := postorder[:len(inLeft)], postorder[len(inLeft):len(postorder)-1]
	root.Left, root.Right = buildTree(inLeft, postLeft), buildTree(inRight, postRight)
	return root
}

func indexOf(s []int, t int) int {
	for i, n := range s {
		if n == t {
			return i
		}
	}
	return -1
}
