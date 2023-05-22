package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var globalHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	globalHead = head
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return buildTree(0, length-1)
}

func buildTree(l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r + 1) / 2
	root := &TreeNode{}
	root.Left = buildTree(l, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree(mid+1, r)
	return root
}
