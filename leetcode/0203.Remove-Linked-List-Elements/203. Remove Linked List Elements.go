package _203_Remove_Linked_List_Elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	p := dummy
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			q := p.Next
			p.Next = q.Next
			q.Next = nil
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}

/*
[1,1]
1
*/
