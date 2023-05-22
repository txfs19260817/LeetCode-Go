package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	p := head
	for i := 0; p != nil && i < m-1; i++ {
		p = p.Next
	}
	if p == nil {
		return head
	}
	q := p.Next
	p.Next = nil
	for i := 0; q != nil && i < n-1; i++ {
		q = q.Next
	}
	if q == nil {
		return head
	}
	nextHead := q.Next
	p.Next = deleteNodes(nextHead, m, n)
	return head
}

/*
[1,2,3,4,5,6,7,8,9,10,11]
1
3
*/
