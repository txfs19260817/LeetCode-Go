package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	front, tail := head, head
	for expLen := 1; front != nil; expLen++ {
		var steps int
		rear, preRear := front, front
		for ; rear != nil && steps < expLen; steps++ {
			preRear = rear
			rear = rear.Next
		}
		if steps%2 == 0 && steps != 0 {
			preRear.Next = nil
			newhead, newtail := reverse(front)
			tail.Next = newhead
			newtail.Next = rear
			preRear = newtail
		}
		front, tail = rear, preRear
	}
	return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	p := head.Next
	head.Next = nil
	tail := head
	for p != nil {
		tmp := p.Next
		p.Next = head
		head = p
		p = tmp
	}
	return head, tail
}
