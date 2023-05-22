package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead, evenHead := head, &ListNode{}
	p := evenHead
	for i := 0; head.Next != nil; i++ {
		if i%2 == 0 { //next is even
			p.Next = head.Next
			head.Next = head.Next.Next
			p = p.Next
			p.Next = nil
		} else {
			head = head.Next
		}
	}

	head.Next = evenHead.Next
	return oddHead
}
