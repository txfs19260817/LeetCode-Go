package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for pre := head; pre != nil; pre = pre.Next {
		cur := pre.Next
		for cur != nil && pre.Val == cur.Val {
			cur = cur.Next
		}
		pre.Next = cur
	}
	return head
}
