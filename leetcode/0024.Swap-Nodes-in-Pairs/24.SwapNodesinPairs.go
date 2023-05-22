package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	head = head.Next
	for p != nil && p.Next != nil {
		q := p.Next // swap 2 nodes
		p.Next = q.Next
		q.Next = p

		p, q = p.Next, q.Next // p: first un-swapped node, q: last swapped node
		if p != nil && p.Next != nil {
			q.Next = p.Next // link q to the next pair's second node which should be the first after swapping
		}
	}
	return head
}
