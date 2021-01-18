package _061_Rotate_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rear, length := head, 1
	for rear.Next != nil {
		rear = rear.Next
		length++
	}
	k %= length

	p, q := head, head
	for ; k > 0; k-- {
		q = q.Next
	}
	for q != rear {
		p, q = p.Next, q.Next
	}
	rear.Next = head
	head = p.Next
	p.Next = nil
	return head
}

/*
[1,2,3,4,5]
2

[0,1,2]
4
*/
