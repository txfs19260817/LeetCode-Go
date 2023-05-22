package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast, length := head, head, 1
	for ; fast.Next != nil; fast = fast.Next {
		length++
	}
	k %= length
	fast = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}
