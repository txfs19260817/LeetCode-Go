package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	for ; n > 0; n-- {
		q = q.Next
	}
	if q == nil {
		return head.Next
	}
	for q.Next != nil {
		p, q = p.Next, q.Next
	}
	p.Next = p.Next.Next
	return head
}

/*
[1,2,3,4,5]
2

[1]
1

[1,2]
1
*/
