package _092_Reverse_Linked_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	dummy := &ListNode{0, head}
	var s, p, q, t *ListNode

	for cur := dummy; cur != nil && n >= -1; m, n, cur = m-1, n-1, cur.Next {
		if m == 1 {
			s = cur
		}
		if m == 0 {
			p = cur
		}
		if n == 0 {
			q = cur
		}
		if n == -1 {
			t = cur
		}
	}
	s.Next, q.Next = nil, nil
	p, q = reverse(p)
	s.Next, q.Next = p, t
	return dummy.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	p, last := head.Next, head
	head.Next = nil
	for p != nil {
		temp := p.Next
		p.Next = head
		head = p
		p = temp
	}
	return head, last
}

/*
[1,2,3,4,5]
2
4

[3,5]
1
2
*/
