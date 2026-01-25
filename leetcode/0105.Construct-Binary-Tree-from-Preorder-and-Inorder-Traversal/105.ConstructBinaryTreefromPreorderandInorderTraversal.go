package leetcode

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	midInorderIdx := slices.Index(inorder, preorder[0])
	inorderLeft, inorderRight := inorder[:midInorderIdx], inorder[midInorderIdx+1:]
	preorderLeft, preorderRight := preorder[1:1+len(inorderLeft)], preorder[1+len(inorderLeft):]
	return &TreeNode{Val: preorder[0], Left: buildTree(preorderLeft, inorderLeft), Right: buildTree(preorderRight, inorderRight)}
}
