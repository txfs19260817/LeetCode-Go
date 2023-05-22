package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	s, f := head, head
	for i := 1; i < k; i++ {
		f = f.Next
	}
	p := f
	for f.Next != nil {
		s, f = s.Next, f.Next
	}
	s.Val, p.Val = p.Val, s.Val
	return head
}
