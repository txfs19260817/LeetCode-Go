package _082_Remove_Duplicates_from_Sorted_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	p, r := head, dummyHead
	for p != nil {
		q := p
		for q != nil && p.Val == q.Val {
			q = q.Next
		}
		if p.Next != q {
			p = q
		} else {
			r.Next = p
			p, r = p.Next, r.Next
			r.Next = nil
		}
	}
	return dummyHead.Next
}

/*
[1,1]
[1,1,1,2,3]
[1,2,3,3,4,4,5]
*/
