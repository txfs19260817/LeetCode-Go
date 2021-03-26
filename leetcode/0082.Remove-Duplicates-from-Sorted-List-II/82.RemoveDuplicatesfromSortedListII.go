package _082_Remove_Duplicates_from_Sorted_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{-114514, head}
	pre := dummyHead
	for cur := head; cur != nil && cur.Next != nil; cur = cur.Next {
		if cur.Val != cur.Next.Val {
			pre = cur
			continue
		}
		next := cur.Next
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		pre.Next = next
		cur = pre
	}
	return dummyHead.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
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
