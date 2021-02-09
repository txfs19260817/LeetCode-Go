package _105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	inRootIdx := indexOf(inorder, preorder[0])
	inLeft, inRight := inorder[:inRootIdx], inorder[inRootIdx+1:]
	preLeft, preRight := preorder[1:1+len(inLeft)], preorder[1+len(inLeft):]
	root.Left, root.Right = buildTree(preLeft, inLeft), buildTree(preRight, inRight)
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
