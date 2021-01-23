package _083_Remove_Duplicates_from_Sorted_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for p := head; p != nil; p = p.Next {
		q := p.Next
		for q != nil && p.Val == q.Val {
			q = q.Next
		}
		p.Next = q
	}
	return head
}

/*
[1,1,2]
[1,1,2,3,3]
*/
