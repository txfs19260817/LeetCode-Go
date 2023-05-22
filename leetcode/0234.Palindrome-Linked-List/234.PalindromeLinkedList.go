package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	temp := middleNode(head)
	mid := temp.Next
	temp.Next = nil

	for mid = reverseList(mid); mid != nil; mid, head = mid.Next, head.Next {
		if mid.Val != head.Val {
			return false
		}
	}
	return true
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
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
