package _143_Reorder_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMid(head)
	temp := mid.Next
	mid.Next = nil
	mid = temp
	mid = reverse(mid)
	p := head
	for mid != nil && p != nil {
		temp = mid.Next
		mid.Next = p.Next
		p.Next = mid
		mid = temp
		p = p.Next.Next
	}
}

func findMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s, f := head, head
	for f.Next != nil && f.Next.Next != nil {
		s, f = s.Next, f.Next.Next
	}
	return s
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head.Next
	head.Next = nil
	for p != nil {
		temp := p.Next
		p.Next = head
		head = p
		p = temp
	}
	return head
}

/*
wrong -> [1,4,3,2]
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	dummyHead := &ListNode{}
	p, r := head, dummyHead
	for p != nil && p.Next != nil { //1->2->3->4
		q := p.Next          //p=1,q=2
		p.Next = p.Next.Next //1->3
		p = p.Next           //p=3
		r.Next = q           //dummy->q=2
		q.Next = nil         //q=2->nil
		r = r.Next           //r=2
	}
	r, p = reverse(dummyHead.Next), head
	for r != nil && p != nil {
		temp := r.Next
		r.Next = p.Next
		p.Next = r
		r = temp
		p = p.Next.Next
	}
}
*/
