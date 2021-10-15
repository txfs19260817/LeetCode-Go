package _008_Construct_Binary_Search_Tree_from_Preorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	firstLargerIndex := -1
	for i, num := range preorder {
		if num > preorder[0] {
			firstLargerIndex = i
			break
		}
	}
	if firstLargerIndex == -1 {
		return &TreeNode{
			Val:  preorder[0],
			Left: bstFromPreorder(preorder[1:]),
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:firstLargerIndex]),
		Right: bstFromPreorder(preorder[firstLargerIndex:]),
	}
}
