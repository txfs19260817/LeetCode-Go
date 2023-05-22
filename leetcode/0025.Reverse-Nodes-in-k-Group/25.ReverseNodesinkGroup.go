package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	nextHead, p := head, head
	for i := 0; i < k; i++ {
		if nextHead == nil {
			return head
		}
		p = nextHead
		nextHead = nextHead.Next
	}

	p.Next = nil
	p = head // last node after reversing
	dummyHead := &ListNode{}
	for head != nil {
		temp := head.Next
		head.Next = dummyHead.Next
		dummyHead.Next = head
		head = temp
	}

	p.Next = reverseKGroup(nextHead, k)
	return dummyHead.Next
}

/*
[1,2,3,4,5]
3

[1,2,3,4,5]
2

[1]
1
*/
