package _083_Remove_Duplicates_from_Sorted_List

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

/*
[1,1,2]
[1,1,2,3,3]
*/
