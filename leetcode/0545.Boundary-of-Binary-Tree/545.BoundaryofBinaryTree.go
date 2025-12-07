package leetcode

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{root.Val}
	if isLeaf(root) {
		return ans
	}

	addLeftBoundary(root, &ans)
	addLeaves(root, &ans)
	addRightBoundary(root, &ans)
	return ans
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}

func addLeftBoundary(root *TreeNode, ans *[]int) {
	for p := root.Left; p != nil; {
		if !isLeaf(p) {
			*ans = append(*ans, p.Val)
		}
		if p.Left != nil {
			p = p.Left
		} else {
			p = p.Right
		}
	}
}

func addLeaves(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	if isLeaf(node) {
		*ans = append(*ans, node.Val)
	}
	addLeaves(node.Left, ans)
	addLeaves(node.Right, ans)
}

func addRightBoundary(root *TreeNode, ans *[]int) {
	var tmp []int
	for p := root.Right; p != nil; {
		if !isLeaf(p) {
			tmp = append(tmp, p.Val)
		}
		if p.Right != nil {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	slices.Reverse(tmp)
	*ans = append(*ans, tmp...)
}
