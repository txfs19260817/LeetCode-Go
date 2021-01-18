package _086_Partition_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	small, large := &ListNode{}, &ListNode{}
	s, l := small, large
	for p := head; p != nil; {
		if p.Val < x {
			s.Next = p
			s, p = s.Next, p.Next
			s.Next = nil
		} else {
			l.Next = p
			l, p = l.Next, p.Next
			l.Next = nil
		}
	}
	s.Next = large.Next
	return small.Next
}
